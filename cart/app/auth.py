from fastapi import Depends, HTTPException, Request, status
import os
import requests
from dotenv import load_dotenv

load_dotenv()
AUTH_SERVICE_URL = os.getenv("AUTH_SERVICE_URL", "http://auth:8002/")

def get_token_from_cookie(request: Request):
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

def get_current_user(token: str = Depends(get_token_from_cookie)):
    cookie = {"access_token": f"{token}"}
    try:
        response = requests.get(f"http://auth:8002/auth/info", cookies=cookie)
        print(response.json())
        if response.status_code != 200:
            raise HTTPException(
                status_code=status.HTTP_401_UNAUTHORIZED,
                detail="Не удалось проверить учетные данные",
                headers={"WWW-Authenticate": "Bearer"},
            )
        user_data = response.json()
        return user_data  # Возвращаем данные пользователя
    except requests.exceptions.ConnectionError:
        raise HTTPException(
            status_code=status.HTTP_503_SERVICE_UNAVAILABLE,
            detail="Сервис аутентификации недоступен",
        )