import uuid

from django.db import models
from django.contrib.auth.models import AbstractUser

class User(AbstractUser):
    ROLE_CHOICES = (
        ('admin', 'Administrator'),
        ('user', 'User'),
    )
    role = models.CharField(max_length=10, choices=ROLE_CHOICES, default='user')

    def __str__(self):
        return self.username

class Product(models.Model):
    name = models.CharField(max_length=255, unique=True, db_index=True)
    short_description = models.TextField()
    full_description = models.TextField()
    composition = models.TextField()
    weight = models.FloatField()
    price = models.DecimalField(max_digits=10, decimal_places=2)
    photo = models.CharField(max_length=255)  # Путь к файлу

    class Meta:
        db_table = 'product'  # Устанавливаем имя таблицы в базе данных
        indexes = [
            models.Index(fields=['name']),  # Индекс для поля name
        ]

    def __str__(self):
        return self.name
