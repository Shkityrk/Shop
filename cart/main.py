# cart/main.py
from fastapi import FastAPI, Depends, HTTPException, status
from sqlalchemy.orm import Session
from typing import List
from app import schemas, models
from app.database import engine, Base, get_db
from app.auth import get_current_user
from fastapi.middleware.cors import CORSMiddleware

Base.metadata.create_all(bind=engine)

app = FastAPI()


# Настройка CORS (если требуется)
origins = [
    "http://localhost",
    "http://auth:8002/",  # или другой адрес фронтенда
]

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# Эндпоинт для получения всех товаров в корзине текущего пользователя
@app.get("/cart/", response_model=List[schemas.CartItem])
def get_cart_items(
    db: Session = Depends(get_db),
    user_data: dict = Depends(get_current_user)
):
    user_id = user_data["id"]
    items = db.query(models.CartItem).filter(models.CartItem.user_id == user_id).all()
    return items


@app.post("/cart/add/", response_model=schemas.CartItem, status_code=status.HTTP_201_CREATED)
def add_item_to_cart(
    item: schemas.CartItemCreate,
    db: Session = Depends(get_db),
    user_data: dict = Depends(get_current_user)
):
    user_id = user_data["id"]
    # Проверяем, есть ли уже такой товар в корзине
    cart_item = db.query(models.CartItem).filter(
        models.CartItem.user_id == user_id,
        models.CartItem.product_id == item.product_id
    ).first()
    if cart_item:
        # Обновляем количество
        cart_item.quantity += item.quantity
    else:
        # Создаем новый элемент корзины
        cart_item = models.CartItem(
            user_id=user_id,
            product_id=item.product_id,
            quantity=item.quantity
        )
        db.add(cart_item)
    db.commit()
    db.refresh(cart_item)
    return cart_item

# Эндпоинт для обновления количества товара в корзине
@app.put("/cart/update/{item_id}", response_model=schemas.CartItem)
def update_cart_item(
    item_id: int,
    item: schemas.CartItemCreate,
    db: Session = Depends(get_db),
    user_data: dict = Depends(get_current_user)
):
    user_id = user_data["id"]
    cart_item = db.query(models.CartItem).filter(
        models.CartItem.id == item_id,
        models.CartItem.user_id == user_id
    ).first()
    if not cart_item:
        raise HTTPException(status_code=404, detail="Элемент корзины не найден")
    cart_item.quantity = item.quantity
    db.commit()
    db.refresh(cart_item)
    return cart_item

# Эндпоинт для удаления товара из корзины
@app.delete("/cart/delete/{item_id}", status_code=status.HTTP_204_NO_CONTENT)
def delete_cart_item(
    item_id: int,
    db: Session = Depends(get_db),
    user_data: dict = Depends(get_current_user)
):
    user_id = user_data["id"]
    cart_item = db.query(models.CartItem).filter(
        models.CartItem.id == item_id,
        models.CartItem.user_id == user_id
    ).first()
    if not cart_item:
        raise HTTPException(status_code=404, detail="Элемент корзины не найден")
    db.delete(cart_item)
    db.commit()
    return
