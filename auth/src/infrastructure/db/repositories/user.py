from sqlalchemy.orm import Session
from src.domain.models.user import User
from src.domain.interfaces.user_repository import AbstractUserRepository
from src.domain.repository import map_user_orm_to_domain
from src.infrastructure.db.models.user_model import UserORM


class UserRepository(AbstractUserRepository):
    def __init__(self, db: Session):
        self.db = db

    def get_by_username(self, username: str) -> User | None:
        orm_user = self.db.query(UserORM).filter(UserORM.username == username).first()
        if orm_user is None:
            return None
        return map_user_orm_to_domain(orm_user)

    def save(self, user: User) -> None:
        orm_user = UserORM(
            id = user.id,
            first_name=user.first_name,
            last_name=user.last_name,
            username=user.username,
            email=user.email,
            hashed_password=user.hashed_password
        )
        self.db.add(orm_user)
        self.db.commit()
        return map_user_orm_to_domain(orm_user)

    def get_by_username_or_email(self, username: str, email: str) -> User | None:
        orm_user = (
            self.db.query(UserORM)
            .filter((UserORM.username == username) | (UserORM.email == email))
            .first()
        )
        if orm_user is None:
            return None
        return map_user_orm_to_domain(orm_user)
