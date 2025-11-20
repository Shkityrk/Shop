"""HTTP реализация CartRepository через Data Service"""
from typing import List, Optional

from src.domain.interfaces.cart_repository import AbstractCartRepository
from src.domain.models import CartItem
from src.infrastructure.http.data_service_client import DataServiceClient


class CartRepositoryHTTP(AbstractCartRepository):
    """Репозиторий корзины через Data Service HTTP API"""
    
    def __init__(self, client: DataServiceClient = None):
        self.client = client or DataServiceClient()
    
    def get_cart_items_by_user_id(self, user_id: int) -> List[CartItem]:
        """Получить все товары в корзине пользователя"""
        try:
            items_data = self.client.get_cart_items_by_user_id(user_id)
            return [self._map_to_domain(item_data) for item_data in items_data]
        except Exception as e:
            # Логируем ошибку и возвращаем пустой список или пробрасываем дальше
            raise Exception(f"Failed to get cart items: {e}")
    
    def get_cart_item_by_id(self, item_id: int, user_id: int) -> Optional[CartItem]:
        """Получить элемент корзины по ID"""
        item_data = self.client.get_cart_item_by_id(item_id)
        if item_data is None:
            return None
        # Проверяем, что элемент принадлежит пользователю
        if item_data.get("user_id") != user_id:
            return None
        return self._map_to_domain(item_data)
    
    def get_cart_item_by_product_id(self, user_id: int, product_id: int) -> Optional[CartItem]:
        """Получить элемент корзины по product_id и user_id"""
        item_data = self.client.get_cart_item_by_user_and_product(user_id, product_id)
        if item_data is None:
            return None
        return self._map_to_domain(item_data)
    
    def add_cart_item(self, cart_item: CartItem) -> CartItem:
        """Добавить новый элемент в корзину"""
        cart_item_data = {
            "user_id": cart_item.user_id,
            "product_id": cart_item.product_id,
            "quantity": cart_item.quantity
        }
        
        created_item = self.client.create_cart_item(cart_item_data)
        return self._map_to_domain(created_item)
    
    def update_cart_item(self, cart_item: CartItem) -> CartItem:
        """Обновить элемент корзины"""
        cart_item_data = {
            "quantity": cart_item.quantity
        }
        
        updated_item = self.client.update_cart_item(cart_item.id, cart_item_data)
        return self._map_to_domain(updated_item)
    
    def delete_cart_item(self, item_id: int, user_id: int) -> None:
        """Удалить элемент из корзины"""
        # Проверяем, что элемент принадлежит пользователю перед удалением
        item = self.get_cart_item_by_id(item_id, user_id)
        if item:
            self.client.delete_cart_item(item_id)
    
    def _map_to_domain(self, item_data: dict) -> CartItem:
        """Преобразовать данные из API в доменную модель"""
        return CartItem(
            id=item_data.get("id"),
            user_id=item_data.get("user_id"),
            product_id=item_data.get("product_id"),
            quantity=item_data.get("quantity"),
        )

