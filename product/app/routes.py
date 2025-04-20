from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy.orm import Session
from . import models, database
from .schemas import ProductCreate

router = APIRouter(prefix="/product")

def get_db():
    db = database.SessionLocal()
    try:
        yield db
    finally:
        db.close()

@router.get("/list")
def read_products(skip: int = 0, limit: int = 100, db: Session = Depends(get_db)):
    products = db.query(models.Product).offset(skip).limit(limit).all()
    return products


@router.post("/add")
def create_product(product: ProductCreate, db: Session = Depends(get_db)):
    # Проверка на уникальность имени
    existing_product = db.query(models.Product).filter(models.Product.name == product.name).first()
    if existing_product:
        raise HTTPException(status_code=400, detail="Product with this name already exists")

    # Создание нового продукта
    new_product = models.Product(
        name=product.name,
        short_description=product.short_description,
        full_description=product.full_description,
        composition=product.composition,
        weight=product.weight,
        price=product.price,
        photo=product.photo
    )
    db.add(new_product)
    db.commit()
    db.refresh(new_product)
    return new_product


@router.put("/update/{id}")
def update_product(id: int, product: ProductCreate, db: Session = Depends(get_db)):
    # Поиск продукта по имени
    db_product = db.query(models.Product).filter(models.Product.id == id).first()
    if not db_product:
        raise HTTPException(status_code=404, detail="Product not found")

    # Обновление полей продукта
    db_product.short_description = product.short_description
    db_product.full_description = product.full_description
    db_product.composition = product.composition
    db_product.weight = product.weight
    db_product.price = product.price
    db_product.photo = product.photo

    db.commit()
    db.refresh(db_product)
    return db_product


@router.get("/verify-product/{product_id}")
def verify_product(product_id: int, db: Session = Depends(get_db)):
    product = db.query(models.Product).filter(models.Product.id == product_id).first()
    if product:
        return {"exists": True}
    else:
        return {"exists": False}

@router.get("/info/{id}")
def info_product(id: str, db: Session = Depends(get_db)):
    product = db.query(models.Product).filter(models.Product.id == id).first()
    if product:
        return {"product": product}
    else:
        raise HTTPException(status_code=404, detail="Product not found")
