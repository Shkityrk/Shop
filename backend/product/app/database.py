from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker

__all__ = [
    "engine",
]

DATABASE_URL = "postgresql://admin_user:admin_pass@db:5432/admin_db"

engine = create_engine(DATABASE_URL)
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)
