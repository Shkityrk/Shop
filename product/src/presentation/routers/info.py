from fastapi import APIRouter, Depends

from src.application.service.product_service import ProductService
from src.presentation.routers.dependences import get_product_service

info_router = APIRouter()

@info_router.get("/info/{id}")
def info_product(id: int,
                 product_service: ProductService = Depends(get_product_service)):
    product = product_service.info_product(id)
    return product