from fastapi import Depends
from sqlalchemy.orm import Session

from src.infrastructure.db.repositories.product import ProductRepository
from src.infrastructure.db.session import get_db
from src.application.service.product_service import ProductService


def get_product_service(db: Session = Depends(get_db)) -> ProductService:
    return ProductService(ProductRepository(db))
