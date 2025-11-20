from fastapi import APIRouter, Depends, Request

from src.presentation.routers.dependences import get_product_service
from src.application.service.product_service import ProductService

list_router = APIRouter()

@list_router.get("/list")
def read_products(skip: int = 0,
                  limit: int = 100,
                  product_service: ProductService = Depends(get_product_service)):
    """Получить список продуктов с пагинацией"""
    products = product_service.read_products(skip=skip, limit=limit)
    return products
