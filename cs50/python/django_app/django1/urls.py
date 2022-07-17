from django.urls import path
from . import views

urlpatterns = [
    path("", views.index, name="index"),
    path("papun",views.papun, name="papun"),
    path("newapp",views.newapp, name="newapp"),
    path("<str:name>",views.greet, name ="greet")


]