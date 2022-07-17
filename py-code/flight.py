class Flight():
    def __init__(self, capacity):
        self.capacity = capacity
        self.passengers = []

    def add_pass(self, name):
        if not self.open():
            return False
        self.passengers.append(name)
        return True
    
    def open(self):
        return self.capacity - len(self.passengers)

flight = Flight(3)

people = ["harry","ron","hermione","ginny"]

for person in people:
    success = flight.add_pass(person)
    if success:
        print(f"added {person} to flight successfully.")
    else:
        print(f"no seats available for {person}")

