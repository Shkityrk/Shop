from sqlalchemy import Column, Integer, ForeignKey
from sqlalchemy.orm import relationship

from src.infrastructure.db.base import Base

__all__ = [
    "InventoryItemORM",
    "InventoryMovementORM",
]


class InventoryItemORM(Base):
    """
    Текущий остаток товара в конкретном bin (ячейке).
    Внешний product_id ссылается на продукт из сервиса product (логически, без FK).
    """

    __tablename__ = "inventory_items"

    id = Column(Integer, primary_key=True, index=True, autoincrement=True)

    product_id = Column(Integer, index=True, nullable=False)
    warehouse_id = Column(Integer, ForeignKey("warehouses.id", ondelete="CASCADE"), nullable=False)
    bin_id = Column(Integer, ForeignKey("bin_locations.id", ondelete="CASCADE"), nullable=False)

    quantity = Column(Integer, nullable=False, default=0)

    warehouse = relationship("WarehouseORM", backref="inventory_items")
    bin = relationship("BinLocationORM", backref="inventory_items")


class InventoryMovementORM(Base):
    """
    История перемещений товара между складами/ячейками.
    """

    __tablename__ = "inventory_movements"

    id = Column(Integer, primary_key=True, index=True, autoincrement=True)

    product_id = Column(Integer, index=True, nullable=False)
    from_warehouse_id = Column(Integer, ForeignKey("warehouses.id"), nullable=True)
    to_warehouse_id = Column(Integer, ForeignKey("warehouses.id"), nullable=True)
    from_bin_id = Column(Integer, ForeignKey("bin_locations.id"), nullable=True)
    to_bin_id = Column(Integer, ForeignKey("bin_locations.id"), nullable=True)

    quantity = Column(Integer, nullable=False)

