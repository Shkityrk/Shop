from pydantic import BaseModel
from typing import Optional, List
from datetime import datetime


class ShipmentItemCreate(BaseModel):
    product_id: int
    quantity: int


class ShipmentCreate(BaseModel):
    order_id: int
    user_id: int
    address: str
    courier_id: Optional[int] = None
    items: List[ShipmentItemCreate]


class ShipmentItemRead(BaseModel):
    id: int
    product_id: int
    quantity: int

    class Config:
        from_attributes = True


class ShipmentRead(BaseModel):
    id: int
    order_id: int
    user_id: int
    address: str
    tracking_code: str
    status: str
    courier_id: Optional[int]
    created_at: datetime
    updated_at: datetime
    items: List[ShipmentItemRead]

    class Config:
        from_attributes = True


class ShipmentStatusUpdate(BaseModel):
    status: str
    courier_id: Optional[int] = None

