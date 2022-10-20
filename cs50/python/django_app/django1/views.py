from django.shortcuts import render
from django.http import HttpResponse

# Create your views here.
def index(request):
    return render(request, "django/index.html")

def papun(request):
    return HttpResponse("Hello papun!")

def newapp(request):
    return HttpResponse("yay! this is another new page!!")

def greet(request, name):
    return render(request, "django/greet.html",{
        "name":name.capitalize()
    })
