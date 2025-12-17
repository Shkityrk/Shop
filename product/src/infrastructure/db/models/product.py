from sqlalchemy import Column, Integer, String, Float, Text, DECIMAL
from src.infrastructure.db.base import Base

__all__ = [
    "ProductORM",
]


class ProductORM(Base):
    __tablename__ = 'product'

    id = Column(Integer,
                primary_key=True,
                index=True,
                autoincrement=True)
    name = Column(String,
                  unique=True,
                  index=True)
    short_description = Column(Text)
    full_description = Column(Text)
    composition = Column(Text)
    weight = Column(Float)
    price = Column(DECIMAL(10, 2))
    photo = Column(String)
    storage_rule_id = Column(Integer, nullable=True)

    def __repr__(self):
        return f"<Product(name={self.name})>"
