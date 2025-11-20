from sqlalchemy import Column, Integer, ForeignKey
from src.infrastructure.db.base import Base

__all__ = [
    "CartItemORM"
]

class CartItemORM(Base):
    __tablename__ = "cart_items"

    id = Column(Integer, primary_key=True, index=True, autoincrement=True)
    user_id = Column(Integer, index=True, nullable=False)
    product_id = Column(Integer, nullable=False)
    quantity = Column(Integer, default=1)

