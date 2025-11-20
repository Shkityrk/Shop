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
    
    # Product methods
    def create_product(self, product_data: dict) -> dict:
        """Создает продукт"""
        return self._request("POST", "/api/products", json=product_data)
    
    def get_product_by_id(self, product_id: int) -> Optional[dict]:
        """Получает продукт по ID"""
        try:
            return self._request("GET", f"/api/products/{product_id}")
        except Exception:
            return None
    
    def get_product_by_name(self, name: str) -> Optional[dict]:
        """Получает продукт по имени"""
        try:
            return self._request("GET", f"/api/products/name/{name}")
        except Exception:
            return None
    
    def list_products(self, skip: int = 0, limit: int = 100) -> List[dict]:
        """Получает список продуктов"""
        params = {}
        if skip > 0:
            params["skip"] = skip
        if limit != 100:
            params["limit"] = limit
        return self._request("GET", "/api/products", params=params)
    
    def update_product(self, product_id: int, product_data: dict) -> dict:
        """Обновляет продукт"""
        return self._request("PUT", f"/api/products/{product_id}", json=product_data)
    
    def delete_product(self, product_id: int) -> None:
        """Удаляет продукт"""
        self._request("DELETE", f"/api/products/{product_id}")
    
    def check_product_exists(self, product_id: int) -> bool:
        """Проверяет существование продукта"""
        try:
            result = self._request("GET", f"/api/products/{product_id}/exists")
            return result.get("exists", False)
        except Exception:
            return False

