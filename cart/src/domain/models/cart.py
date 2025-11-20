from dataclasses import dataclass
from typing import Optional


@dataclass
class CartItem:
    user_id: int
    product_id: int
    quantity: int
    id: Optional[int] = None

