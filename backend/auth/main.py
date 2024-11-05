from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware

from app.database import engine
from app.models import Base
from app.EnvArrayProcessor import EnvArrayProcessor
from app.routes import auth_router

# Создаем таблицы
Base.metadata.create_all(bind=engine)

app = FastAPI(root_path="/auth")

# Инициализируем список разрешенных источников
processor = EnvArrayProcessor("../.hosts")
origins = processor.get_unique_urls()
print(origins)

# Добавляем CORS middleware
app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# Подключаем маршруты
app.include_router(auth_router)

# Путь для проверки здоровья сервиса
@app.get("/api/health")
def read_health():
    return {"status": "ok"}
