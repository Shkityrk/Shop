from fastapi import FastAPI, Depends, HTTPException, Response, Request
from sqlalchemy.orm import Session

from . import models, schemas
from .EnvArrayProcessor import EnvArrayProcessor
from .database import engine, SessionLocal
from .utils import hash_password
from .auth import (
    authenticate_user,
    create_access_token,
    get_current_user,
    ACCESS_TOKEN_EXPIRE_MINUTES,
)
from datetime import timedelta

from fastapi.middleware.cors import CORSMiddleware

models.Base.metadata.create_all(bind=engine)

app = FastAPI()

processor = EnvArrayProcessor("../.hosts")
origins = processor.get_unique_urls()

print(origins)

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,  # Чтобы разрешить куки
    allow_methods=["*"],
    allow_headers=["*"],
)


# Зависимость для получения сессии базы данных
def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()


@app.post("/register", response_model=schemas.UserOut)
def register(
    user_create: schemas.UserCreate, response: Response, db: Session = Depends(get_db)
):
    user = (
        db.query(models.User)
        .filter(
            models.User.username
            == user_create.username | models.User.email
            == user_create.email
        )
        .first()
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

    # Перенаправляем на нужную страницу (например, на главную)
    return new_user


@app.post("/login")
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
    # Устанавливаем токен в куки
    response.set_cookie(
        key="access_token", value=f"Bearer {access_token}", httponly=False
    )
    return {"message": "Успешный вход"}


@app.post("/logout")
def logout(response: Response):
    # Удаляем токен из куки
    response.delete_cookie(key="access_token")
    return {"message": "Успешный выход"}


@app.get("/protected-route")
def protected_route(current_user: models.User = Depends(get_current_user)):
    return {"message": f"Привет, {current_user.username}!"}


# Обработчик для проверки токена из куки
@app.middleware("http")
async def check_jwt_token(request: Request, call_next):
    try:
        token = request.cookies.get("access_token")
        if token:
            scheme, _, param = token.partition(" ")
            request.headers["Authorization"] = f"{scheme} {param}"
    except Exception:
        pass
    response = await call_next(request)
    return response


@app.get("/api/health")
def read_health():
    return {"status": "ok"}
