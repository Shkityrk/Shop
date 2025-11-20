from fastapi import Depends

from src.infrastructure.http.product_repository_http import ProductRepositoryHTTP
from src.application.service.product_service import ProductService


def get_product_service() -> ProductService:
    """
    Возвращает ProductService с HTTP репозиторием для Data Service.
    Теперь мы не обращаемся к БД напрямую, а используем Data Service API.
    """
    return ProductService(ProductRepositoryHTTP())
