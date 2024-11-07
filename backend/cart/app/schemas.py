from pydantic import BaseModel

class CartItemBase(BaseModel):
    product_id: int
    quantity: int = 1

class CartItemCreate(CartItemBase):
    pass

class CartItem(CartItemBase):
    id: int
    user_id: int

    class Config:
        orm_mode = True

class UserOut(BaseModel):
    id: int
    username: str
    email: str
    first_name: str
    last_name: str

    class Config:
        orm_mode = True