from fastapi import APIRouter, Depends, Request, status

from src.presentation.routers.dependences import get_cart_service, get_user_id
from src.application.services.cart_service import CartService
from src.schemas.schemas import CartItem, CartItemCreate

add_router = APIRouter()


@add_router.post("/add", response_model=CartItem, status_code=status.HTTP_201_CREATED)
def add_item_to_cart(
    item: CartItemCreate,
    request: Request,
    user_id: int = Depends(get_user_id),
    cart_service: CartService = Depends(get_cart_service)
):
    """
    Добавить товар в корзину.
    User ID извлекается из заголовка X-User-Id, установленного Gateway middleware.
    """
    cart_item = cart_service.add_item_to_cart(user_id, item)
    return cart_item

