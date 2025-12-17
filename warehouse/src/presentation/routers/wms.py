from fastapi import APIRouter, Depends
from fastapi import HTTPException
from sqlalchemy import func
from sqlalchemy.orm import Session

from src.infrastructure.db.session import get_db
from src.infrastructure.db.models.inventory import InventoryItemORM, InventoryMovementORM
from src.schemas.schemas import (
    WmsCheckRequest,
    WmsCheckResponse,
    WmsCommitResponse,
    WmsShortage,
    WmsCommitResultItem,
    WmsAllocation,
)

wms_router = APIRouter(prefix="/wms", tags=["wms"])


@wms_router.post("/check", response_model=WmsCheckResponse)
def check_availability(
    data: WmsCheckRequest,
    db: Session = Depends(get_db),
):
    shortages: list[WmsShortage] = []

    for item in data.items:
        total = (
            db.query(func.coalesce(func.sum(InventoryItemORM.quantity), 0))
            .filter(InventoryItemORM.product_id == item.product_id)
            .scalar()
        )
        if total < item.quantity:
            shortages.append(
                WmsShortage(
                    product_id=item.product_id,
                    requested=item.quantity,
                    available=total,
                )
            )

    return WmsCheckResponse(ok=len(shortages) == 0, shortages=shortages)


@wms_router.post("/commit", response_model=WmsCommitResponse)
def commit_order(
    data: WmsCheckRequest,
    db: Session = Depends(get_db),
):
    shortages: list[WmsShortage] = []

    # Сначала проверяем наличие, чтобы не начинать транзакцию впустую
    for item in data.items:
        total = (
            db.query(func.coalesce(func.sum(InventoryItemORM.quantity), 0))
            .filter(InventoryItemORM.product_id == item.product_id)
            .scalar()
        )
        if total < item.quantity:
            shortages.append(
                WmsShortage(
                    product_id=item.product_id,
                    requested=item.quantity,
                    available=total,
                )
            )

    if shortages:
        return WmsCommitResponse(ok=False, shortages=shortages, items=[])

    result_items: list[WmsCommitResultItem] = []

    try:
        for item in data.items:
            remaining = item.quantity
            allocations: list[WmsAllocation] = []

            inv_rows = (
                db.query(InventoryItemORM)
                .filter(InventoryItemORM.product_id == item.product_id)
                .order_by(InventoryItemORM.quantity.desc())
                .with_for_update()
                .all()
            )

            for row in inv_rows:
                if remaining <= 0:
                    break
                deduct = min(row.quantity, remaining)
                if deduct <= 0:
                    continue

                row.quantity -= deduct
                remaining -= deduct

                allocations.append(
                    WmsAllocation(
                        warehouse_id=row.warehouse_id,
                        bin_id=row.bin_id,
                        deducted=deduct,
                    )
                )

                movement = InventoryMovementORM(
                    product_id=item.product_id,
                    from_warehouse_id=row.warehouse_id,
                    to_warehouse_id=None,
                    from_bin_id=row.bin_id,
                    to_bin_id=None,
                    quantity=deduct,
                )
                db.add(movement)

            if remaining > 0:
                # Это не должно случиться после предварительной проверки, но на всякий случай
                db.rollback()
                raise HTTPException(
                    status_code=400,
                    detail=f"Not enough stock for product {item.product_id}",
                )

            result_items.append(
                WmsCommitResultItem(
                    product_id=item.product_id,
                    deducted=item.quantity,
                    allocations=allocations,
                )
            )

        db.commit()
        return WmsCommitResponse(ok=True, shortages=[], items=result_items)
    except HTTPException:
        # HTTPException пробрасываем дальше
        raise
    except Exception as exc:  # noqa: BLE001
        # Любая другая ошибка = откат и 500
        db.rollback()
        raise HTTPException(status_code=500, detail=str(exc)) from exc

