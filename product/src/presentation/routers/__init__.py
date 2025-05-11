from fastapi import APIRouter
from .list import list_router
from .add import add_router
from .update import update_router
from .verify import verify_router
from .info import info_router
from .healthcheck import healthcheck_router

__all__ = [
    "root_router"
]

root_router = APIRouter(prefix="/product")

root_router.include_router(info_router)
root_router.include_router(list_router)
root_router.include_router(add_router)
root_router.include_router(verify_router)
root_router.include_router(healthcheck_router)
root_router.include_router(update_router)
