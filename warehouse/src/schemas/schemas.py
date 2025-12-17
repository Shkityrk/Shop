from pydantic import BaseModel
from typing import Optional


class WarehouseBase(BaseModel):
    name: str
    address: str
    working_hours: Optional[str] = None


class WarehouseCreate(WarehouseBase):
    pass


class WarehouseRead(WarehouseBase):
    id: int

    class Config:
        from_attributes = True


class StorageRuleBase(BaseModel):
    name: str
    description: Optional[str] = None
    is_hazardous: bool = False
    is_oversized: bool = False
    temp_min: Optional[float] = None
    temp_max: Optional[float] = None


class StorageRuleCreate(StorageRuleBase):
    pass


class StorageRuleRead(StorageRuleBase):
    id: int

    class Config:
        from_attributes = True


class BinLocationBase(BaseModel):
    warehouse_id: int
    zone: str
    aisle: str
    rack: str
    bin_code: str
    storage_rule_id: Optional[int] = None


class BinLocationCreate(BinLocationBase):
    # Опционально: сразу добавить продукт на полку
    product_id: Optional[int] = None
    quantity: Optional[int] = None


class BinLocationRead(BinLocationBase):
    id: int
    # Информация о товаре на полке (если есть)
    product_id: Optional[int] = None
    quantity: Optional[int] = None

    class Config:
        from_attributes = True


class InventoryItemBase(BaseModel):
    product_id: int
    warehouse_id: int
    bin_id: int
    quantity: int


class InventoryItemCreate(InventoryItemBase):
    pass


class InventoryItemRead(InventoryItemBase):
    id: int

    class Config:
        from_attributes = True


class InventoryMoveRequest(BaseModel):
    product_id: int
    from_warehouse_id: int
    from_bin_id: int
    to_warehouse_id: int
    to_bin_id: int
    quantity: int


class ProductTotalQuantity(BaseModel):
    product_id: int
    total_quantity: int


class WmsItem(BaseModel):
    product_id: int
    quantity: int


class WmsCheckRequest(BaseModel):
    items: list[WmsItem]


class WmsShortage(BaseModel):
    product_id: int
    requested: int
    available: int


class WmsAllocation(BaseModel):
    warehouse_id: int
    bin_id: int
    deducted: int


class WmsCommitResultItem(BaseModel):
    product_id: int
    deducted: int
    allocations: list[WmsAllocation]


class WmsCheckResponse(BaseModel):
    ok: bool
    shortages: list[WmsShortage]


class WmsCommitResponse(BaseModel):
    ok: bool
    shortages: list[WmsShortage]
    items: list[WmsCommitResultItem]

