from uuid import uuid4
from typing import List
from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy.orm import Session
import random
import string

from src.infrastructure.db.session import get_db
from src.infrastructure.db.models.shipment import ShipmentORM, ShipmentItemORM
from src.schemas.schemas import (
    ShipmentCreate,
    ShipmentRead,
    ShipmentStatusUpdate,
)

shipping_router = APIRouter(tags=["shipping"])


def _generate_tracking_code(db: Session, prefix: str = "SH", digits: int = 9, attempts: int = 10) -> str:
    """Generate a readable tracking code and ensure it's unique in the DB.

    Format: <prefix><digits> e.g. SH123456789
    Attempts up to `attempts` times to avoid collisions.
    """
    for _ in range(attempts):
        code = prefix + "".join(random.choices(string.digits, k=digits))
        # quick uniqueness check
        if not db.query(ShipmentORM).filter(ShipmentORM.tracking_code == code).first():
            return code
    raise HTTPException(status_code=500, detail="Failed to generate unique tracking code")


@shipping_router.get("/list", response_model=List[ShipmentRead])
def list_shipments(
    skip: int = 0,
    limit: int = 100,
    status: str | None = None,
    db: Session = Depends(get_db),
):
    """Get a list of all shipments, optionally filtered by status."""
    query = db.query(ShipmentORM)
    if status:
        query = query.filter(ShipmentORM.status == status)
    shipments = query.order_by(ShipmentORM.created_at.desc()).offset(skip).limit(limit).all()
    # подгрузим items для каждой доставки
    for s in shipments:
        s.items  # noqa: B018
    return shipments


@shipping_router.post("/", response_model=ShipmentRead)
def create_shipment(
    data: ShipmentCreate,
    db: Session = Depends(get_db),
):
    # generate readable, trackable and unique tracking code
    tracking_code = _generate_tracking_code(db)

    shipment = ShipmentORM(
        order_id=data.order_id,
        user_id=data.user_id,
        address=data.address,
        tracking_code=tracking_code,
        status="created",
        courier_id=data.courier_id,
    )
    db.add(shipment)
    db.flush()  # получаем id для FK

    for item in data.items:
        db.add(
            ShipmentItemORM(
                shipment_id=shipment.id,
                product_id=item.product_id,
                quantity=item.quantity,
            )
        )

    db.commit()
    db.refresh(shipment)
    # подгрузим items для ответа
    shipment.items  # noqa: B018
    return shipment


@shipping_router.get("/{tracking_code}", response_model=ShipmentRead)
def get_shipment(tracking_code: str, db: Session = Depends(get_db)):
    shipment = (
        db.query(ShipmentORM)
        .filter(ShipmentORM.tracking_code == tracking_code)
        .first()
    )
    if not shipment:
        raise HTTPException(status_code=404, detail="Shipment not found")
    shipment.items  # noqa: B018
    return shipment


@shipping_router.patch("/{tracking_code}/status", response_model=ShipmentRead)
def update_shipment_status(
    tracking_code: str,
    data: ShipmentStatusUpdate,
    db: Session = Depends(get_db),
):
    shipment = (
        db.query(ShipmentORM)
        .filter(ShipmentORM.tracking_code == tracking_code)
        .first()
    )
    if not shipment:
        raise HTTPException(status_code=404, detail="Shipment not found")

    shipment.status = data.status
    if data.courier_id is not None:
        shipment.courier_id = data.courier_id

    db.commit()
    db.refresh(shipment)
    shipment.items  # noqa: B018
    return shipment
