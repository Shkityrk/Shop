from pydantic import BaseModel

class UserCreate(BaseModel):
    username: str
    email: str
    password: str
    first_name: str
    last_name: str

class UserResponse(BaseModel):
    id: int
    username: str
    email: str
    first_name: str
    last_name: str
    role: str

    class Config:
        orm_mode = True

class TokenData(BaseModel):
    username: str = None
