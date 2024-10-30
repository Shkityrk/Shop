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
    name = models.CharField(max_length=255, unique=True)
    short_description = models.TextField(max_length=255)
    full_description = models.TextField(max_length=255)
    composition = models.TextField(max_length=255)
    weight = models.FloatField(max_length=255)
    price = models.DecimalField(max_digits=10, decimal_places=10)

    def __str__(self) ->str:
        return self.name
