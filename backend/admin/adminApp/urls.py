from django.urls import path, include
from rest_framework import routers
from .views import UserViewSet, ProductViewSet

router = routers.DefaultRouter()
router.register(r'users', UserViewSet)
router.register(r'products', ProductViewSet)

urlpatterns = [
    path('api/', include(router.urls)),
]
