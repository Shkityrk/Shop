from src.domain.models import CartItem
from src.infrastructure.db.models.cart_model import CartItemORM


def map_cart_item_orm_to_domain(orm_item: CartItemORM) -> CartItem:
    return CartItem(
        id=orm_item.id,
        user_id=orm_item.user_id,
        product_id=orm_item.product_id,
        quantity=orm_item.quantity,
    )

