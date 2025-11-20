from fastapi import APIRouter, Depends, Request, status

from src.presentation.routers.dependences import get_cart_service, get_user_id
from src.application.services.cart_service import CartService

delete_router = APIRouter()


@delete_router.delete("/delete/{item_id}", status_code=status.HTTP_204_NO_CONTENT)
def delete_cart_item(
    item_id: int,
    request: Request,
    user_id: int = Depends(get_user_id),
    cart_service: CartService = Depends(get_cart_service)
):
    """
    Удалить товар из корзины.
    User ID извлекается из заголовка X-User-Id, установленного Gateway middleware.
    """
    cart_service.delete_cart_item(user_id, item_id)
    return

