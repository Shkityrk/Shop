from sqlalchemy import Column, Integer, String, Float, Text, DECIMAL
from sqlalchemy.ext.declarative import declarative_base

Base = declarative_base()

class Product(Base):
    __tablename__ = 'product'

    id = Column(Integer, primary_key=True, index=True)
    name = Column(String, unique=True, index=True)
    short_description = Column(Text)
    full_description = Column(Text)
    composition = Column(Text)
    weight = Column(Float)
    price = Column(DECIMAL(10, 2))
    photo = Column(String)  # Путь к файлу

    def __repr__(self):
        return f"<Product(name={self.name})>"
