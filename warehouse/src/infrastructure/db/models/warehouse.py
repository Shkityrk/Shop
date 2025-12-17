from sqlalchemy import Column, Integer, String, Text
from src.infrastructure.db.base import Base

__all__ = [
    "WarehouseORM",
]


class WarehouseORM(Base):
    __tablename__ = "warehouses"

    id = Column(Integer, primary_key=True, index=True, autoincrement=True)
    name = Column(String, unique=True, index=True, nullable=False)
    address = Column(Text, nullable=False)
    working_hours = Column(String, nullable=True)

    def __repr__(self) -> str:
        return f"<Warehouse(id={self.id}, name={self.name})>"

