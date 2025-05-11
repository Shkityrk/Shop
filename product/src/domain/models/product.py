from dataclasses import dataclass
from typing import Optional

from src.infrastructure.db.models.product import ProductORM


@dataclass
class Product:
    name: str
    short_description: str
    full_description: str
    composition: str
    weight: float
    price: float
    photo: str
    id: Optional[int] = None

    @classmethod
    def from_orm(cls, orm_prod: ProductORM) -> "Product":
        return cls(
            id = orm_prod.id,
            name = orm_prod.name,
            short_description = orm_prod.short_description,
            full_description = orm_prod.full_description,
            composition = orm_prod.composition,
            weight = orm_prod.weight,
            price = orm_prod.price,
            photo = orm_prod.photo,
        )