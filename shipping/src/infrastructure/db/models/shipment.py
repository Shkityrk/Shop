from datetime import datetime

from sqlalchemy import Column, Integer, BigInteger, String, Text, DateTime, ForeignKey
from sqlalchemy.orm import relationship

from src.infrastructure.db.base import Base

__all__ = [
    "ShipmentORM",
    "ShipmentItemORM",
]


class ShipmentORM(Base):
    __tablename__ = "shipments"

    id = Column(Integer, primary_key=True, autoincrement=True, index=True)
    order_id = Column(BigInteger, index=True, nullable=False)
    user_id = Column(Integer, index=True, nullable=False)
    address = Column(Text, nullable=False)
    tracking_code = Column(String, unique=True, index=True, nullable=False)
    status = Column(String, nullable=False, default="created")
    courier_id = Column(Integer, nullable=True, index=True)
    created_at = Column(DateTime, default=datetime.utcnow, nullable=False)
    updated_at = Column(
        DateTime,
        default=datetime.utcnow,
        onupdate=datetime.utcnow,
        nullable=False,
    )

    items = relationship("ShipmentItemORM", back_populates="shipment", cascade="all, delete-orphan")


class ShipmentItemORM(Base):
    __tablename__ = "shipment_items"

    id = Column(Integer, primary_key=True, autoincrement=True, index=True)
    shipment_id = Column(Integer, ForeignKey("shipments.id", ondelete="CASCADE"), nullable=False)
    product_id = Column(Integer, nullable=False)
    quantity = Column(Integer, nullable=False)

    shipment = relationship("ShipmentORM", back_populates="items")

