from pydantic import BaseModel
from typing import Optional


class ProductCreate(BaseModel):
    name: str
    short_description: str
    full_description: str
    composition: str
    weight: float
    price: float
    photo: str
    storage_rule_id: Optional[int] = None
