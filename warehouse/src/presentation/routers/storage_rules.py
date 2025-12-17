from typing import List

from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy.orm import Session

from src.infrastructure.db.session import get_db
from src.infrastructure.db.models.storage_rule import StorageRuleORM
from src.schemas.schemas import StorageRuleCreate, StorageRuleRead

storage_rule_router = APIRouter(prefix="/storage-rules", tags=["storage_rules"])


@storage_rule_router.post("/", response_model=StorageRuleRead)
def create_storage_rule(
    data: StorageRuleCreate,
    db: Session = Depends(get_db),
):
    rule = StorageRuleORM(**data.model_dump())
    db.add(rule)
    db.commit()
    db.refresh(rule)
    return rule


@storage_rule_router.get("/", response_model=List[StorageRuleRead])
def list_storage_rules(db: Session = Depends(get_db)):
    rules = db.query(StorageRuleORM).all()
    return rules


@storage_rule_router.delete("/{rule_id}")
def delete_storage_rule(rule_id: int, db: Session = Depends(get_db)):
    rule = db.query(StorageRuleORM).filter(StorageRuleORM.id == rule_id).first()
    if not rule:
        raise HTTPException(status_code=404, detail="Storage rule not found")
    db.delete(rule)
    db.commit()
    return {"status": "deleted"}

