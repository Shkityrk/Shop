from fastapi import APIRouter, Depends, Request

from src.presentation.routers.dependences import get_cart_service, get_user_id
from src.application.services.cart_service import CartService
from src.schemas.schemas import CartItem, CartItemCreate

update_router = APIRouter()


@update_router.put("/update/{item_id}", response_model=CartItem)
def update_cart_item(
    item_id: int,
    item: CartItemCreate,
    request: Request,
    user_id: int = Depends(get_user_id),
    cart_service: CartService = Depends(get_cart_service)
):
    """
    Обновить количество товара в корзине.
    User ID извлекается из заголовка X-User-Id, установленного Gateway middleware.
    """
    cart_item = cart_service.update_cart_item(user_id, item_id, item)
    return cart_item

