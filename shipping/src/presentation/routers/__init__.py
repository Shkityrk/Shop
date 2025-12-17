from fastapi import APIRouter

from .shipping import shipping_router
from .healthcheck import healthcheck_router

__all__ = [
    "root_router",
]

root_router = APIRouter(prefix="/shipping")

root_router.include_router(healthcheck_router)
root_router.include_router(shipping_router)

