# alembic/env.py
from logging.config import fileConfig
from sqlalchemy import engine_from_config
from sqlalchemy import pool
from alembic import context
import sys
import os
from dotenv import load_dotenv

from app import Base

# Добавьте путь к проекту

sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), '..')))

# Загрузите переменные окружения из .env
load_dotenv()

# Получение конфигурации Alembic
config = context.config

# Настройка логирования
fileConfig(config.config_file_name)

# Импортируйте ваши модели

target_metadata = Base.metadata

# Получите URL подключения из переменной окружения
database_url = os.getenv("DATABASE_URL")
if not database_url:
    raise ValueError("Нет переменной окружения DATABASE_URL. Проверьте ваш .env файл.")

# Обновите конфигурацию Alembic с новым URL
config.set_main_option("sqlalchemy.url", database_url)

def run_migrations_offline():
    """Run migrations in 'offline' mode."""
    url = config.get_main_option("sqlalchemy.url")
    context.configure(
        url=url, target_metadata=target_metadata, literal_binds=True
    )

    with context.begin_transaction():
        context.run_migrations()

def run_migrations_online():
    """Run migrations in 'online' mode."""
    connectable = engine_from_config(
        config.get_section(config.config_ini_section),
        prefix='sqlalchemy.',
        poolclass=pool.NullPool,
    )

    with connectable.connect() as connection:
        context.configure(
            connection=connection, target_metadata=target_metadata
        )

        with context.begin_transaction():
            context.run_migrations()

if context.is_offline_mode():
    run_migrations_offline()
else:
    run_migrations_online()
