from fastapi import FastAPI
from app.routes import router
from fastapi.middleware.cors import CORSMiddleware
from app.database import engine, Base
from app.models import Product


Base.metadata.create_all(bind=engine)
app = FastAPI()

# Настройка CORS (если требуется)
origins = [
    "http://localhost",
    "http://auth:8002/",  # или другой адрес фронтенда
    "http://localhost:5173",

]


# Добавляем CORS middleware
app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

app.include_router(router)
