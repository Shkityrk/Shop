from typing import List, Optional

from sqlalchemy.orm import Session

from src.domain.interfaces.product_repository import AbstractProductRepository
from src.domain.models import Product
from src.domain.repo import map_product_orm_to_domain
from src.infrastructure.db.models.product import ProductORM


class ProductRepository(AbstractProductRepository):
    def __init__(self, db: Session):
        self.db = db

    def get_all_products(self, skip: int = 0, limit: int = 100) -> List[Product]:
        db_items = (
            self.db
            .query(ProductORM)
            .offset(skip)
            .limit(limit)
            .all()
        )
        return [self._to_domain(item) for item in db_items]

    def get_product_by_id(self, product_id: int) -> Optional[Product]:
        """
        Найти продукт по ID
        """
        db_item = (
            self.db
            .query(ProductORM)
            .filter(ProductORM.id == product_id)
            .first()
        )
        return self._to_domain(db_item) if db_item else None

    def get_by_name(self, name: str) -> Optional[Product]:
        """
        Найти продукт по имени
        """
        db_item = (
            self.db
            .query(ProductORM)
            .filter(ProductORM.name == name)
            .first()
        )
        return self._to_domain(db_item) if db_item else None

    def check_if_exists(self, name: str) -> bool:
        """
        Проверить, существует ли продукт с таким именем
        """
        return self.db.query(ProductORM).filter(ProductORM.name == name).first() is not None

    def add_product(self, product: Product) -> Product:
        """
        Создать новый продукт
        """
        db_item = self._to_entity(product)
        self.db.add(db_item)
        self.db.commit()
        self.db.refresh(db_item)
        return self._to_domain(db_item)

    def save(self, product: Product) -> Product:
        """
        Создать или обновить продукт. Если в доменном объекте есть id — обновляем, иначе — создаём.
        """
        if product.id:
            db_item = (
                self.db
                .query(ProductORM)
                .filter(ProductORM.id == product.id)
                .first()
            )
            if not db_item:
                raise ValueError(f"Product with id={product.id} not found")
            # обновляем поля
            db_item.name = product.name
            db_item.short_description = product.short_description
            db_item.full_description = product.full_description
            db_item.composition = product.composition
            db_item.weight = product.weight
            db_item.price = product.price
            db_item.photo = product.photo
            self.db.commit()
            self.db.refresh(db_item)
            return self._to_domain(db_item)
        else:
            return self.add_product(product)

    def _to_domain(self, db_item: ProductORM) -> Product:
        """
        Преобразовать ORM-модель в доменный объект
        """
        return Product(
            id=db_item.id,
            name=db_item.name,
            short_description=db_item.short_description,
            full_description=db_item.full_description,
            composition=db_item.composition,
            weight=db_item.weight,
            price=float(db_item.price),
            photo=db_item.photo,
        )

    def _to_entity(self, product: Product) -> ProductORM:
        """
        Преобразовать доменный объект в ORM-модель
        """
        return ProductORM(
            name=product.name,
            short_description=product.short_description,
            full_description=product.full_description,
            composition=product.composition,
            weight=product.weight,
            price=product.price,
            photo=product.photo,
        )
