from fastapi import APIRouter, Depends, HTTPException, Response
from sqlalchemy.orm import Session
from datetime import timedelta

from src.schemas import schemas
from src.domain import models
from src.infrastructure.database import SessionLocal
from src.domain.utils import hash_password
from src.domain.auth import (
    authenticate_user,
    create_access_token,
    get_current_user,
    ACCESS_TOKEN_EXPIRE_MINUTES,
)

auth_router = APIRouter(prefix="/auth")



# Зависимость для получения сессии базы данных
def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

# Регистрация
@auth_router.post("/register", response_model=schemas.UserOut)
def register(
    user_create: schemas.UserCreate, response: Response, db: Session = Depends(get_db)
):
    user = (
        db.query(models.User).filter(models.User.username == user_create.username or
                                     models.User.email == user_create.email).first()
    )
    if user:
        raise HTTPException(
            status_code=400,
            detail="Пользователь с таким именем или email уже существует",
        )
    hashed_password = hash_password(user_create.password)
    new_user = models.User(
        first_name=user_create.first_name,
        last_name=user_create.last_name,
        username=user_create.username,
        email=user_create.email,
        hashed_password=hashed_password,
    )
    db.add(new_user)
    db.commit()
    db.refresh(new_user)
    access_token_expires = timedelta(minutes=ACCESS_TOKEN_EXPIRE_MINUTES)
    access_token = create_access_token(
        data={"sub": new_user.username}, expires_delta=access_token_expires
    )
    response.set_cookie(
        key="access_token", value=f"Bearer {access_token}", httponly=False
    )
    return new_user

# Вход
@auth_router.post("/login")
def login(
    response: Response, user_login: schemas.UserLogin, db: Session = Depends(get_db)
):
    user = authenticate_user(db, user_login.username, user_login.password)
    if not user:
        raise HTTPException(
            status_code=401, detail="Неверное имя пользователя или пароль"
        )
    access_token_expires = timedelta(minutes=ACCESS_TOKEN_EXPIRE_MINUTES)
    access_token = create_access_token(
        data={"sub": user.username}, expires_delta=access_token_expires
    )
    response.set_cookie(
        key="access_token",
        value=f"Bearer {access_token}",  # ваш токен
        httponly=True,
    )
    response.set_cookie(
        key="access_token",
        value=f"Bearer {access_token}",  # ваш токен
        httponly=True,
        domain="localhost",
            )

    return {"message": "Успешный вход"}

# Выход
@auth_router.post("/logout")
def logout(response: Response):
    response.delete_cookie(key="access_token")
    return {"message": "Успешный выход"}

# Защищенный маршрут
@auth_router.get("/protected-route")
def protected_route(current_user: models.User = Depends(get_current_user)):
    return {"message": f"Привет, {current_user.username}!"}



@auth_router.get("/info", response_model=schemas.UserOut)
def verify_token(current_user: models.User = Depends(get_current_user)):
    return current_user