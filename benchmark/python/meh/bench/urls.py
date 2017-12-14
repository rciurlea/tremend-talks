from django.conf.urls import url
from .views import hello, stats

urlpatterns = [
    url(r'hello', hello, name='hello'),
    url(r'stats', stats, name='stats'),
]
