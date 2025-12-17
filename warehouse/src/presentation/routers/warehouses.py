from typing import List

from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy.orm import Session

from src.infrastructure.db.session import get_db
from src.infrastructure.db.models.warehouse import WarehouseORM
from src.schemas.schemas import WarehouseCreate, WarehouseRead

warehouse_router = APIRouter(prefix="/warehouses", tags=["warehouses"])


@warehouse_router.post("/", response_model=WarehouseRead)
def create_warehouse(
    data: WarehouseCreate,
    db: Session = Depends(get_db),
):
    warehouse = WarehouseORM(**data.model_dump())
    db.add(warehouse)
    db.commit()
    db.refresh(warehouse)
    return warehouse


@warehouse_router.get("/", response_model=List[WarehouseRead])
def list_warehouses(db: Session = Depends(get_db)):
    warehouses = db.query(WarehouseORM).all()
    return warehouses


@warehouse_router.get("/{warehouse_id}", response_model=WarehouseRead)
def get_warehouse(warehouse_id: int, db: Session = Depends(get_db)):
    warehouse = db.query(WarehouseORM).filter(WarehouseORM.id == warehouse_id).first()
    if not warehouse:
        raise HTTPException(status_code=404, detail="Warehouse not found")
    return warehouse


@warehouse_router.delete("/{warehouse_id}")
def delete_warehouse(warehouse_id: int, db: Session = Depends(get_db)):
    warehouse = db.query(WarehouseORM).filter(WarehouseORM.id == warehouse_id).first()
    if not warehouse:
        raise HTTPException(status_code=404, detail="Warehouse not found")
    db.delete(warehouse)
    db.commit()
    return {"status": "deleted"}

