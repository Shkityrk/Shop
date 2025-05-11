from fastapi import APIRouter, Depends, Request

from src.presentation.routers.dependences import get_product_service
from src.application.service.product_service import ProductService
from src.schemas.schemas import ProductCreate

add_router = APIRouter()


@add_router.post("/add")
def add_product(product: ProductCreate,
                product_service: ProductService = Depends(get_product_service)
                ):
    product = product_service.create_product(product)
    return product
