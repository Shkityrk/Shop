"""HTTP клиент для работы с Data Service"""
import requests
from typing import Optional, List

from src.config import DATA_SERVICE_URL


class DataServiceClient:
    """Клиент для обращения к Data Service API"""
    
    def __init__(self):
        self.base_url = DATA_SERVICE_URL
        self.session = requests.Session()
        self.session.headers.update({"Content-Type": "application/json"})
    
    def _request(self, method: str, endpoint: str, **kwargs):
        """Выполняет HTTP запрос к Data Service"""
        url = f"{self.base_url}{endpoint}"
        
        try:
            response = self.session.request(method, url, timeout=5, **kwargs)
            response.raise_for_status()
            
            if response.status_code == 204:
                return None
            return response.json()
        except requests.exceptions.ConnectionError as e:
            raise Exception(f"Data Service is unavailable: {e}")
        except requests.exceptions.Timeout as e:
            raise Exception(f"Data Service request timeout: {e}")
        except requests.exceptions.RequestException as e:
            raise Exception(f"Data Service request failed: {e}")
    
    # Cart methods
    def create_cart_item(self, cart_item_data: dict) -> dict:
        """Создает элемент корзины"""
        return self._request("POST", "/api/cart", json=cart_item_data)
    
    def get_cart_item_by_id(self, item_id: int) -> Optional[dict]:
        """Получает элемент корзины по ID"""
        try:
            return self._request("GET", f"/api/cart/{item_id}")
        except Exception:
            return None
    
    def get_cart_items_by_user_id(self, user_id: int) -> List[dict]:
        """Получает корзину пользователя"""
        try:
            return self._request("GET", f"/api/cart/user/{user_id}")
        except Exception:
            return []
    
    def get_cart_item_by_user_and_product(self, user_id: int, product_id: int) -> Optional[dict]:
        """Получает элемент корзины по user_id и product_id"""
        try:
            return self._request("GET", f"/api/cart/user/{user_id}/product/{product_id}")
        except Exception:
            return None
    
    def update_cart_item(self, item_id: int, cart_item_data: dict) -> dict:
        """Обновляет элемент корзины"""
        return self._request("PUT", f"/api/cart/{item_id}", json=cart_item_data)
    
    def delete_cart_item(self, item_id: int) -> None:
        """Удаляет элемент корзины"""
        self._request("DELETE", f"/api/cart/{item_id}")

