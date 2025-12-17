from fastapi import APIRouter

from .warehouses import warehouse_router
from .locations import location_router
from .storage_rules import storage_rule_router
from .inventory import inventory_router
from .healthcheck import healthcheck_router
from .wms import wms_router

__all__ = [
    "root_router",
]

root_router = APIRouter(prefix="/warehouse")

root_router.include_router(healthcheck_router)
root_router.include_router(warehouse_router)
root_router.include_router(location_router)
root_router.include_router(storage_rule_router)
root_router.include_router(inventory_router)
root_router.include_router(wms_router)

