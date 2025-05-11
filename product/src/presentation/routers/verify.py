from fastapi import APIRouter, Depends, Request

from src.presentation.routers.dependences import get_product_service
from src.application.service.product_service import ProductService
from src.schemas.schemas import ProductCreate

verify_router = APIRouter()

@verify_router.get("/verify/{name}")
def add_product(name: str,
                product_service: ProductService = Depends(get_product_service)
                ):
    product = product_service.verify_product(name)
    return product
