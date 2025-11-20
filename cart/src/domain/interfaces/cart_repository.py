from abc import ABC, abstractmethod
from typing import List, Optional

from src.domain.models import CartItem


class AbstractCartRepository(ABC):

    @abstractmethod
    def get_cart_items_by_user_id(self, user_id: int) -> List[CartItem]:
        pass

    @abstractmethod
    def get_cart_item_by_id(self, item_id: int, user_id: int) -> Optional[CartItem]:
        pass

    @abstractmethod
    def get_cart_item_by_product_id(self, user_id: int, product_id: int) -> Optional[CartItem]:
        pass

    @abstractmethod
    def add_cart_item(self, cart_item: CartItem) -> CartItem:
        pass

    @abstractmethod
    def update_cart_item(self, cart_item: CartItem) -> CartItem:
        pass

    @abstractmethod
    def delete_cart_item(self, item_id: int, user_id: int) -> None:
        pass

