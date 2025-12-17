from typing import List

from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy.orm import Session

from src.infrastructure.db.session import get_db
from src.infrastructure.db.models.location import BinLocationORM
from src.infrastructure.db.models.inventory import InventoryItemORM, InventoryMovementORM
from src.infrastructure.db.models.warehouse import WarehouseORM
from src.schemas.schemas import BinLocationCreate, BinLocationRead

location_router = APIRouter(prefix="/locations", tags=["locations"])


@location_router.post("/bins", response_model=BinLocationRead)
def create_bin_location(
    data: BinLocationCreate,
    db: Session = Depends(get_db),
):
    warehouse = db.query(WarehouseORM).filter(WarehouseORM.id == data.warehouse_id).first()
    if not warehouse:
        raise HTTPException(status_code=404, detail="Warehouse not found")

    # Создаём ячейку (полку)
    bin_location = BinLocationORM(
        warehouse_id=data.warehouse_id,
        zone=data.zone,
        aisle=data.aisle,
        rack=data.rack,
        bin_code=data.bin_code,
        storage_rule_id=data.storage_rule_id,
    )
    db.add(bin_location)
    db.flush()  # получаем id для FK

    product_id = None
    quantity = None

    # Если указан product_id и quantity — создаём inventory item
    if data.product_id is not None and data.quantity is not None and data.quantity > 0:
        inventory_item = InventoryItemORM(
            product_id=data.product_id,
            warehouse_id=data.warehouse_id,
            bin_id=bin_location.id,
            quantity=data.quantity,
        )
        db.add(inventory_item)

        # Записываем движение товара (поступление)
        movement = InventoryMovementORM(
            product_id=data.product_id,
            from_warehouse_id=None,
            to_warehouse_id=data.warehouse_id,
            from_bin_id=None,
            to_bin_id=bin_location.id,
            quantity=data.quantity,
        )
        db.add(movement)

        product_id = data.product_id
        quantity = data.quantity

    db.commit()
    db.refresh(bin_location)

    # Формируем ответ с информацией о товаре
    return BinLocationRead(
        id=bin_location.id,
        warehouse_id=bin_location.warehouse_id,
        zone=bin_location.zone,
        aisle=bin_location.aisle,
        rack=bin_location.rack,
        bin_code=bin_location.bin_code,
        storage_rule_id=bin_location.storage_rule_id,
        product_id=product_id,
        quantity=quantity,
    )


@location_router.get("/bins", response_model=List[BinLocationRead])
def list_bin_locations(
    warehouse_id: int | None = None,
    db: Session = Depends(get_db),
):
    query = db.query(BinLocationORM)
    if warehouse_id is not None:
        query = query.filter(BinLocationORM.warehouse_id == warehouse_id)

    bins = query.all()
    result = []
    for b in bins:
        # Получаем inventory item для этой ячейки (если есть)
        inventory = db.query(InventoryItemORM).filter(InventoryItemORM.bin_id == b.id).first()
        result.append(BinLocationRead(
            id=b.id,
            warehouse_id=b.warehouse_id,
            zone=b.zone,
            aisle=b.aisle,
            rack=b.rack,
            bin_code=b.bin_code,
            storage_rule_id=b.storage_rule_id,
            product_id=inventory.product_id if inventory else None,
            quantity=inventory.quantity if inventory else None,
        ))
    return result


@location_router.delete("/bins/{bin_id}")
def delete_bin_location(bin_id: int, db: Session = Depends(get_db)):
    bin_location = db.query(BinLocationORM).filter(BinLocationORM.id == bin_id).first()
    if not bin_location:
        raise HTTPException(status_code=404, detail="Bin not found")
    db.delete(bin_location)
    db.commit()
    return {"status": "deleted"}

