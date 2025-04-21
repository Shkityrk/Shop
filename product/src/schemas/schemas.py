from pydantic import BaseModel


class ProductCreate(BaseModel):
    name: str
    short_description: str
    full_description: str
    composition: str
    weight: float
    price: float
    photo: str