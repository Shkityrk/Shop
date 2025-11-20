from fastapi import APIRouter, Depends, Request
from typing import List

from .add import add_router
from .update import update_router
from .delete import delete_router
from .healthcheck import healthcheck_router
from .dependences import get_cart_service, get_user_id
from src.application.services.cart_service import CartService
from src.schemas.schemas import CartItem

__all__ = [
    "root_router"
]

root_router = APIRouter(prefix="/cart")


@root_router.get("", response_model=List[CartItem])
def get_cart_items(
    request: Request,
    user_id: int = Depends(get_user_id),
    cart_service: CartService = Depends(get_cart_service)
):
    """
    Получить все товары в корзине пользователя.
    User ID извлекается из заголовка X-User-Id, установленного Gateway middleware.
    """
    items = cart_service.get_cart_items(user_id)
    return items


root_router.include_router(add_router)
root_router.include_router(update_router)
root_router.include_router(delete_router)
root_router.include_router(healthcheck_router)

