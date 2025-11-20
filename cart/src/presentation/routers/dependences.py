from fastapi import Depends, HTTPException, Request, status

from src.infrastructure.http.cart_repository_http import CartRepositoryHTTP
from src.application.services.cart_service import CartService


def get_cart_service() -> CartService:
    """
    Возвращает CartService с HTTP репозиторием для Data Service.
    Теперь мы не обращаемся к БД напрямую, а используем Data Service API.
    """
    return CartService(CartRepositoryHTTP())


def get_user_id(request: Request) -> int:
    """
    Извлекает user_id из заголовка X-User-Id, который устанавливается Gateway middleware.
    Gateway проверяет токен и добавляет user_id в заголовки для downstream сервисов.
    """
    user_id_header = request.headers.get("X-User-Id")
    if not user_id_header:
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail="User ID not found in headers. Authentication required.",
        )
    
    try:
        user_id = int(user_id_header)
        return user_id
    except ValueError:
        raise HTTPException(
            status_code=status.HTTP_400_BAD_REQUEST,
            detail="Invalid user ID format",
        )

