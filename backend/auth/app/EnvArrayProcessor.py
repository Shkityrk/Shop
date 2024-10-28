import os
from dotenv import load_dotenv

__all__=[
    "EnvArrayProcessor"
]

class EnvArrayProcessor:
    def __init__(self, env_file: str):
        # Загружаем .env файл
        load_dotenv(env_file)
        # Извлекаем массив строк из переменной окружения
        self.url_list = self._parse_env_array(os.getenv("URL_LIST", ""))

    def _parse_env_array(self, env_value: str):
        """
        Преобразует строку из .env файла в список URL.
        Строка может быть представлена в формате:
        ["http://localhost:3000", "http://localhost:8084", ...]
        """
        if not env_value:
            return []
        # Убираем лишние символы и
        # разделяем строки по запятым
        return [url.strip().strip('"') for url in env_value.strip("[]").split(",")]

    def get_unique_urls(self):
        """Возвращает уникальные URL из списка"""
        return list(set(self.url_list))

    def filter_urls_by_port(self, port: int):
        """Фильтрует URL по порту"""
        return [url for url in self.url_list if f":{port}" in url]

    def get_urls_by_domain(self, domain: str):
        """Возвращает URL, содержащие указанный домен"""
        return [url for url in self.url_list if domain in url]

    def print_urls(self):
        """Выводит все URL на экран"""
        for url in self.url_list:
            print(url)


# Пример использования
if __name__ == "__main__":
    processor = EnvArrayProcessor("../.hosts")
    print("Все URL:")
    processor.print_urls()
    print("\nУникальные URL:")
    print(processor.get_unique_urls())
    print("\nФильтр по порту 8000:")
    print(processor.filter_urls_by_port(8000))
    print("\nURL, содержащие 'localhost':")
    print(processor.get_urls_by_domain("localhost"))
