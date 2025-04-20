from fastapi import Depends
from sqlalchemy.orm import Session

from src.infrastructure.db.repositories.user import UserRepository
from src.infrastructure.db.session import get_db
from src.application.services.auth_service import AuthService



def get_auth_service(db: Session = Depends(get_db)) -> AuthService:
    return AuthService(UserRepository(db))