FROM python:3.9-slim


# Установить рабочую директорию
WORKDIR /app

# Установить зависимости системы
RUN apt-get update && apt-get install -y \
    build-essential \
    && rm -rf /var/lib/apt/lists/*

# Копировать файлы требований
COPY auth/requirements.txt .

# Установить Python зависимости
RUN pip install --no-cache-dir -r requirements.txt

# Копировать весь код приложения
COPY ./auth .

# Команда запуска приложения
CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "8000"]
