class Cars:
    def __init__(self):
        self.company = "Toyota"
        self.name = "supra"
        self.colour = "black"
        self.price = 2000000
    def buy(self):
        return f"The {self.company} {self.name} you want to buy is {self.price} and is of {self.colour} colour."
cars = Cars()

print(cars.buy())