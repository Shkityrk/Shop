from typing import List
from fastapi import HTTPException, status

from src.domain.interfaces.cart_repository import AbstractCartRepository
from src.domain.models import CartItem
from src.schemas.schemas import CartItemCreate


class CartService:
    def __init__(self, cart_repo: AbstractCartRepository):
        self.cart_repo = cart_repo

    def get_cart_items(self, user_id: int) -> List[CartItem]:
        """Получить все товары в корзине пользователя"""
        return self.cart_repo.get_cart_items_by_user_id(user_id)

    def add_item_to_cart(self, user_id: int, item: CartItemCreate) -> CartItem:
        """Добавить товар в корзину или увеличить количество, если товар уже есть"""
        existing_item = self.cart_repo.get_cart_item_by_product_id(user_id, item.product_id)
        
        if existing_item:
            # Увеличиваем количество существующего товара
            existing_item.quantity += item.quantity
            return self.cart_repo.update_cart_item(existing_item)
        else:
            # Создаем новый элемент корзины
            new_item = CartItem(
                user_id=user_id,
                product_id=item.product_id,
                quantity=item.quantity
            )
            return self.cart_repo.add_cart_item(new_item)

    def update_cart_item(self, user_id: int, item_id: int, item: CartItemCreate) -> CartItem:
        """Обновить количество товара в корзине"""
        cart_item = self.cart_repo.get_cart_item_by_id(item_id, user_id)
        if not cart_item:
            raise HTTPException(
                status_code=status.HTTP_404_NOT_FOUND,
                detail="Элемент корзины не найден"
            )
        
        cart_item.quantity = item.quantity
        return self.cart_repo.update_cart_item(cart_item)

    def delete_cart_item(self, user_id: int, item_id: int) -> None:
        """Удалить товар из корзины"""
        cart_item = self.cart_repo.get_cart_item_by_id(item_id, user_id)
        if not cart_item:
            raise HTTPException(
                status_code=status.HTTP_404_NOT_FOUND,
                detail="Элемент корзины не найден"
            )
        
        self.cart_repo.delete_cart_item(item_id, user_id)

