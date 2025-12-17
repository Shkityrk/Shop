from sqlalchemy import Column, Integer, String, ForeignKey
from sqlalchemy.orm import relationship

from src.infrastructure.db.base import Base

__all__ = [
    "BinLocationORM",
]


class BinLocationORM(Base):
    """
    Одна запись = один bin (ячейка хранения) внутри конкретного склада.
    Уровни вложенности кодируются полями zone / aisle / rack / bin.
    """

    __tablename__ = "bin_locations"

    id = Column(Integer, primary_key=True, index=True, autoincrement=True)
    warehouse_id = Column(Integer, ForeignKey("warehouses.id", ondelete="CASCADE"), nullable=False)

    zone = Column(String, nullable=False)
    aisle = Column(String, nullable=False)
    rack = Column(String, nullable=False)
    bin_code = Column(String, nullable=False)

    storage_rule_id = Column(Integer, ForeignKey("storage_rules.id"), nullable=True)

    warehouse = relationship("WarehouseORM", backref="bins")

