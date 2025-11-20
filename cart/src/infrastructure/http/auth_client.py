from fastapi import Depends, HTTPException, Request, status
import requests

from src.config import AUTH_SERVICE_URL


def get_token_from_cookie(request: Request) -> str:
    """Извлечь токен из cookie"""
    token = request.cookies.get("access_token")
    if not token:
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail="Токен не найден",
            headers={"WWW-Authenticate": "Bearer"},
        )
    if token.startswith("Bearer "):
        token = token[7:]
    return token


def get_current_user(token: str = Depends(get_token_from_cookie)) -> dict:
    """Получить данные текущего пользователя из auth сервиса"""
    cookie = {"access_token": f"{token}"}
    try:
        response = requests.get(f"{AUTH_SERVICE_URL}/auth/info", cookies=cookie)
        if response.status_code != 200:
            raise HTTPException(
                status_code=status.HTTP_401_UNAUTHORIZED,
                detail="Не удалось проверить учетные данные",
                headers={"WWW-Authenticate": "Bearer"},
            )
        user_data = response.json()
        return user_data
    except requests.exceptions.ConnectionError:
        raise HTTPException(
            status_code=status.HTTP_503_SERVICE_UNAVAILABLE,
            detail="Сервис аутентификации недоступен",
        )

