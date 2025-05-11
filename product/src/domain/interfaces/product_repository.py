from abc import ABC, abstractmethod
from typing import List

from src.domain.models import Product

class AbstractProductRepository(ABC):

    @abstractmethod
    def get_all_products(self) -> List[Product]:
        pass

    @abstractmethod
    def get_info_products(self) -> List[Product]:
        pass

    @abstractmethod
    def get_product_by_id(self, product_id: int) -> Product:
        pass

    @abstractmethod
    def save(self, product: Product):
        pass

