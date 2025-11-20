from fastapi import APIRouter, Depends, Request

from src.presentation.routers.dependences import get_product_service
from src.application.service.product_service import ProductService
from src.schemas.schemas import ProductCreate

update_router = APIRouter()

@update_router.put("/update/{id}")
def update_product(id: int,
                   product: ProductCreate,
                   product_service: ProductService = Depends(get_product_service)):
    """
    Обновить продукт.
    TODO: сделать возможным изменение нескольких полей, а не отправлять все поля сразу
    """
    updated_product = product_service.update_product(id, product)
    return updated_product
