from fastapi import HTTPException

from src.domain.interfaces.product_repository import AbstractProductRepository
from src.domain.models import Product
from src.infrastructure.db.models.product import ProductORM
from src.infrastructure.db.repositories.product import ProductRepository
from src.schemas.schemas import ProductCreate


class ProductService:
    def __init__(self,
                 product_repo: AbstractProductRepository):
        self.product_repo = product_repo

    def read_products(self,
                      skip: int = 0,
                      limit: int = 100,
                      ):
        products = self.product_repo.get_all_products(skip, limit)

        return products

    def create_product(self,
                       product: ProductCreate
                       ) -> None | Product:
        product_exists = self.product_repo.check_if_exists(product.name)
        if product_exists:
            raise HTTPException(status_code=400, detail="Product with this name already exists")
        else:
            new_product = Product(
                name=product.name,
                short_description=product.short_description,
                full_description=product.full_description,
                composition=product.composition,
                weight=product.weight,
                price=product.price,
                photo=product.photo
            )

            return self.product_repo.add_product(new_product)

    def update_product(self,
                       id: int,
                       product: ProductCreate
                       ):
        product_exists = self.product_repo.check_if_exists(product.name)
        if product_exists:
            raise HTTPException(status_code=400, detail="Product with this name already exists")
        else:
            self.product_repo.save(product)

    def verify_product(self,
                       name: str
                       ):
        product_exists = self.product_repo.get_by_name(name)
        if product_exists:
            return {"exists": True,
                    "id": product_exists.id,
                    "name": product_exists.name}
        else:
            return {"exists": False,
                    "id": None,
                    "name": None}

    def info_product(self,
                     id: int
                     ):
        product_exists = self.product_repo.get_product_by_id(id)
        if product_exists:
            return {"product": product_exists}
        else:
            raise HTTPException(status_code=404, detail="Product not found")