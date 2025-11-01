from fastapi import Depends

from src.infrastructure.http.user_repository_http import UserRepositoryHTTP
from src.application.services.auth_service import AuthService


def get_auth_service() -> AuthService:
    """
    Возвращает AuthService с HTTP репозиторием для Data Service.
    Теперь мы не обращаемся к БД напрямую, а используем Data Service API.
    """
    return AuthService(UserRepositoryHTTP())