from dataclasses import dataclass
from typing import Optional

from src.infrastructure.db.models.user_model import UserORM


@dataclass
class User:
    first_name: str
    last_name: str
    username: str
    email: str
    hashed_password: str
    id: Optional[int] = None

    @classmethod
    def from_orm(cls, orm_user: UserORM) -> "User":
        return cls(
            id=orm_user.id,
            first_name=orm_user.first_name,
            last_name=orm_user.last_name,
            username=orm_user.username,
            email=orm_user.email,
            hashed_password=orm_user.hashed_password
        )