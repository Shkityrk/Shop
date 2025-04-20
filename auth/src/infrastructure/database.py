from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

from src.config import (
    DATABASE_USER,
    DATABASE_HOST,
    DATABASE_PASSWORD,
    DATABASE_PORT,
    DATABASE_NAME

)

import os

__all__=[
    "Base",
    "SessionLocal"
]

DATABASE_URL = f"postgresql://{DATABASE_USER}:{DATABASE_PASSWORD}@{DATABASE_HOST}:{DATABASE_PORT}/{DATABASE_NAME}"

try:
    engine = create_engine(DATABASE_URL)
    connection = engine.connect()
    connection.close()

    SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)

except Exception as e:
    print(f"Ошибка подключения к базе данных: {e}")

Base = declarative_base()
