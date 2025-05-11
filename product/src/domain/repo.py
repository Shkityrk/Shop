from src.domain.models import Product
from src.infrastructure.db.models.product import ProductORM


def map_product_orm_to_domain(orm_prod: ProductORM) -> Product:
    return Product(
        id=orm_prod.id,
        name=orm_prod.name,
        short_description=orm_prod.short_description,
        full_description=orm_prod.full_description,
        composition=orm_prod.composition,
        weight=orm_prod.weight,
        price=orm_prod.price,
        photo=orm_prod.photo
    )