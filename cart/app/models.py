# models.py
from sqlalchemy import Column, Integer, ForeignKey
from .database import Base

class CartItem(Base):
    __tablename__ = "cart_items"

    id = Column(Integer, primary_key=True, index=True)
    user_id = Column(Integer, index=True, nullable=False)
    product_id = Column(Integer, nullable=False)
    quantity = Column(Integer, default=1)
