from abc import ABC, abstractmethod
from typing import List, Optional

from src.domain.models import Product


class AbstractProductRepository(ABC):

    @abstractmethod
    def get_all_products(self, skip, limit) -> List[Product]:
        pass

    @abstractmethod
    def get_product_by_id(self, product_id: int) -> Product:
        pass

    @abstractmethod
    def get_by_name(self, name: str) -> Optional[Product]:
        pass

    @abstractmethod
    def check_if_exists(self, name: str) -> bool:
        pass

    @abstractmethod
    def add_product(self, product: Product) -> Product:
        pass

    @abstractmethod
    def save(self, product: Product):
        pass
