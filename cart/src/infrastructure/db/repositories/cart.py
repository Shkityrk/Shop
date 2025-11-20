from typing import List, Optional
from sqlalchemy.orm import Session

from src.domain.interfaces.cart_repository import AbstractCartRepository
from src.domain.models import CartItem
from src.domain.repository import map_cart_item_orm_to_domain
from src.infrastructure.db.models.cart_model import CartItemORM


class CartRepository(AbstractCartRepository):
    def __init__(self, db: Session):
        self.db = db

    def get_cart_items_by_user_id(self, user_id: int) -> List[CartItem]:
        """Получить все товары в корзине пользователя"""
        db_items = (
            self.db
            .query(CartItemORM)
            .filter(CartItemORM.user_id == user_id)
            .all()
        )
        return [map_cart_item_orm_to_domain(item) for item in db_items]

    def get_cart_item_by_id(self, item_id: int, user_id: int) -> Optional[CartItem]:
        """Получить элемент корзины по ID"""
        db_item = (
            self.db
            .query(CartItemORM)
            .filter(
                CartItemORM.id == item_id,
                CartItemORM.user_id == user_id
            )
            .first()
        )
        return map_cart_item_orm_to_domain(db_item) if db_item else None

    def get_cart_item_by_product_id(self, user_id: int, product_id: int) -> Optional[CartItem]:
        """Получить элемент корзины по product_id и user_id"""
        db_item = (
            self.db
            .query(CartItemORM)
            .filter(
                CartItemORM.user_id == user_id,
                CartItemORM.product_id == product_id
            )
            .first()
        )
        return map_cart_item_orm_to_domain(db_item) if db_item else None

    def add_cart_item(self, cart_item: CartItem) -> CartItem:
        """Добавить новый элемент в корзину"""
        db_item = CartItemORM(
            user_id=cart_item.user_id,
            product_id=cart_item.product_id,
            quantity=cart_item.quantity
        )
        self.db.add(db_item)
        self.db.commit()
        self.db.refresh(db_item)
        return map_cart_item_orm_to_domain(db_item)

    def update_cart_item(self, cart_item: CartItem) -> CartItem:
        """Обновить элемент корзины"""
        db_item = (
            self.db
            .query(CartItemORM)
            .filter(CartItemORM.id == cart_item.id)
            .first()
        )
        if not db_item:
            raise ValueError(f"Cart item with id={cart_item.id} not found")
        
        db_item.quantity = cart_item.quantity
        self.db.commit()
        self.db.refresh(db_item)
        return map_cart_item_orm_to_domain(db_item)

    def delete_cart_item(self, item_id: int, user_id: int) -> None:
        """Удалить элемент из корзины"""
        db_item = (
            self.db
            .query(CartItemORM)
            .filter(
                CartItemORM.id == item_id,
                CartItemORM.user_id == user_id
            )
            .first()
        )
        if db_item:
            self.db.delete(db_item)
            self.db.commit()

