from pydantic import BaseModel, EmailStr


__all__=[
    "UserLogin",
    "UserBase",
    "UserOut",
    "UserCreate"
]


class UserBase(BaseModel):
    first_name: str
    last_name: str
    username: str
    email: EmailStr


class UserCreate(UserBase):
    password: str


class UserLogin(BaseModel):
    username: str
    password: str


class UserOut(UserBase):
    id: int

    class Config:
        orm_mode = True
