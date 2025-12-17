from sqlalchemy import Column, Integer, String, Boolean, Float

from src.infrastructure.db.base import Base

__all__ = [
    "StorageRuleORM",
]


class StorageRuleORM(Base):
    __tablename__ = "storage_rules"

    id = Column(Integer, primary_key=True, index=True, autoincrement=True)
    name = Column(String, unique=True, index=True, nullable=False)
    description = Column(String, nullable=True)

    is_hazardous = Column(Boolean, default=False)
    is_oversized = Column(Boolean, default=False)
    temp_min = Column(Float, nullable=True)
    temp_max = Column(Float, nullable=True)

