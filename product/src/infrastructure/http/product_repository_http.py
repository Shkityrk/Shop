"""HTTP реализация ProductRepository через Data Service"""
from typing import List, Optional

from src.domain.interfaces.product_repository import AbstractProductRepository
from src.domain.models import Product
from src.infrastructure.http.data_service_client import DataServiceClient


class ProductRepositoryHTTP(AbstractProductRepository):
    """Репозиторий продуктов через Data Service HTTP API"""
    
    def __init__(self, client: DataServiceClient = None):
        self.client = client or DataServiceClient()
    
    def get_all_products(self, skip: int = 0, limit: int = 100) -> List[Product]:
        """Получить все продукты"""
        products_data = self.client.list_products(skip=skip, limit=limit)
        return [self._map_to_domain(product_data) for product_data in products_data]
    
    def get_product_by_id(self, product_id: int) -> Optional[Product]:
        """Получить продукт по ID"""
        product_data = self.client.get_product_by_id(product_id)
        if product_data is None:
            return None
        return self._map_to_domain(product_data)
    
    def get_by_name(self, name: str) -> Optional[Product]:
        """Получить продукт по имени"""
        product_data = self.client.get_product_by_name(name)
        if product_data is None:
            return None
        return self._map_to_domain(product_data)
    
    def check_if_exists(self, name: str) -> bool:
        """Проверить существование продукта по имени"""
        product = self.get_by_name(name)
        return product is not None
    
    def add_product(self, product: Product) -> Product:
        """Создать новый продукт"""
        product_data = {
            "name": product.name,
            "short_description": product.short_description,
            "full_description": product.full_description,
            "composition": product.composition,
            "weight": product.weight,
            "price": product.price,
            "photo": product.photo
        }
        
        created_product = self.client.create_product(product_data)
        return self._map_to_domain(created_product)
    
    def save(self, product: Product) -> Product:
        """Создать или обновить продукт"""
        if product.id:
            # Обновление существующего продукта
            product_data = {
                "name": product.name,
                "short_description": product.short_description,
                "full_description": product.full_description,
                "composition": product.composition,
                "weight": product.weight,
                "price": product.price,
                "photo": product.photo
            }
            updated_product = self.client.update_product(product.id, product_data)
            return self._map_to_domain(updated_product)
        else:
            # Создание нового продукта
            return self.add_product(product)
    
    def _map_to_domain(self, product_data: dict) -> Product:
        """Преобразовать данные из API в доменную модель"""
        return Product(
            id=product_data.get("id"),
            name=product_data.get("name"),
            short_description=product_data.get("short_description"),
            full_description=product_data.get("full_description"),
            composition=product_data.get("composition"),
            weight=product_data.get("weight"),
            price=product_data.get("price"),
            photo=product_data.get("photo"),
        )

