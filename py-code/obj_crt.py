class Creation:
    def __new__(self):
        return []

myobj = Creation()

print(type(myobj))