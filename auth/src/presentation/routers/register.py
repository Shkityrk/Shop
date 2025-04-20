from fastapi import APIRouter, Depends, Response

from src.presentation.routers.dependences import get_auth_service
from src.schemas import schemas
from src.application.services.auth_service import AuthService

register_router = APIRouter()

@register_router.post("/register", response_model=schemas.UserOut)
def register(
    user_create: schemas.UserCreate,
    response: Response,
    auth_service: AuthService = Depends(get_auth_service)
):
    user = auth_service.register_user(user_create)
    access_token = auth_service.create_access_token(data={"sub": user.username})
    response.set_cookie(
        key="access_token", value=f"Bearer {access_token}", httponly=True
    )
    return user