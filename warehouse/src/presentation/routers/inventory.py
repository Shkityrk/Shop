from typing import List

from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy import func
from sqlalchemy.orm import Session

from src.infrastructure.db.session import get_db
from src.infrastructure.db.models.inventory import InventoryItemORM, InventoryMovementORM
from src.infrastructure.db.models.location import BinLocationORM
from src.infrastructure.db.models.warehouse import WarehouseORM
from src.schemas.schemas import (
    InventoryItemCreate,
    InventoryItemRead,
    InventoryMoveRequest,
    ProductTotalQuantity,
)

inventory_router = APIRouter(prefix="/inventory", tags=["inventory"])


@inventory_router.post("/add", response_model=InventoryItemRead)
def add_inventory_item(
    data: InventoryItemCreate,
    db: Session = Depends(get_db),
):
    warehouse = db.query(WarehouseORM).filter(WarehouseORM.id == data.warehouse_id).first()
    if not warehouse:
        raise HTTPException(status_code=404, detail="Warehouse not found")

    bin_location = db.query(BinLocationORM).filter(BinLocationORM.id == data.bin_id).first()
    if not bin_location:
        raise HTTPException(status_code=404, detail="Bin not found")

    item = (
        db.query(InventoryItemORM)
        .filter(
            InventoryItemORM.product_id == data.product_id,
            InventoryItemORM.bin_id == data.bin_id,
        )
        .first()
    )
    if item:
        item.quantity += data.quantity
    else:
        item = InventoryItemORM(**data.model_dump())
        db.add(item)

    movement = InventoryMovementORM(
        product_id=data.product_id,
        from_warehouse_id=None,
        to_warehouse_id=data.warehouse_id,
        from_bin_id=None,
        to_bin_id=data.bin_id,
        quantity=data.quantity,
    )
    db.add(movement)

    db.commit()
    db.refresh(item)
    return item


@inventory_router.post("/move")
def move_inventory(
    data: InventoryMoveRequest,
    db: Session = Depends(get_db),
):
    # Проверяем склад/ячейки
    from_wh = db.query(WarehouseORM).filter(WarehouseORM.id == data.from_warehouse_id).first()
    to_wh = db.query(WarehouseORM).filter(WarehouseORM.id == data.to_warehouse_id).first()
    if not from_wh or not to_wh:
        raise HTTPException(status_code=404, detail="Warehouse not found")

    from_bin = db.query(BinLocationORM).filter(BinLocationORM.id == data.from_bin_id).first()
    to_bin = db.query(BinLocationORM).filter(BinLocationORM.id == data.to_bin_id).first()
    if not from_bin or not to_bin:
        raise HTTPException(status_code=404, detail="Bin not found")

    # Источник
    from_item = (
        db.query(InventoryItemORM)
        .filter(
            InventoryItemORM.product_id == data.product_id,
            InventoryItemORM.bin_id == data.from_bin_id,
        )
        .first()
    )
    if not from_item or from_item.quantity < data.quantity:
        raise HTTPException(status_code=400, detail="Not enough quantity in source bin")

    # Получатель
    to_item = (
        db.query(InventoryItemORM)
        .filter(
            InventoryItemORM.product_id == data.product_id,
            InventoryItemORM.bin_id == data.to_bin_id,
        )
        .first()
    )
    if to_item:
        to_item.quantity += data.quantity
    else:
        to_item = InventoryItemORM(
            product_id=data.product_id,
            warehouse_id=data.to_warehouse_id,
            bin_id=data.to_bin_id,
            quantity=data.quantity,
        )
        db.add(to_item)

    from_item.quantity -= data.quantity

    movement = InventoryMovementORM(
        product_id=data.product_id,
        from_warehouse_id=data.from_warehouse_id,
        to_warehouse_id=data.to_warehouse_id,
        from_bin_id=data.from_bin_id,
        to_bin_id=data.to_bin_id,
        quantity=data.quantity,
    )
    db.add(movement)

    db.commit()

    return {"status": "moved"}


@inventory_router.get("/items", response_model=List[InventoryItemRead])
def list_inventory_items(
    product_id: int | None = None,
    warehouse_id: int | None = None,
    db: Session = Depends(get_db),
):
    query = db.query(InventoryItemORM)
    if product_id is not None:
        query = query.filter(InventoryItemORM.product_id == product_id)
    if warehouse_id is not None:
        query = query.filter(InventoryItemORM.warehouse_id == warehouse_id)
    return query.all()


@inventory_router.get("/product/{product_id}/total", response_model=ProductTotalQuantity)
def get_product_total_quantity(
    product_id: int,
    db: Session = Depends(get_db),
):
    total = (
        db.query(func.coalesce(func.sum(InventoryItemORM.quantity), 0))
        .filter(InventoryItemORM.product_id == product_id)
        .scalar()
    )
    return ProductTotalQuantity(product_id=product_id, total_quantity=total)


@inventory_router.get("/totals", response_model=List[ProductTotalQuantity])
def get_all_products_totals(db: Session = Depends(get_db)):
    rows = (
        db.query(
            InventoryItemORM.product_id,
            func.sum(InventoryItemORM.quantity).label("total_quantity"),
        )
        .group_by(InventoryItemORM.product_id)
        .all()
    )
    return [
        ProductTotalQuantity(product_id=row.product_id, total_quantity=row.total_quantity)
        for row in rows
    ]

