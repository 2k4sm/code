# Learning OOP in Python.

class Employee:
    
    def __init__(self,firstname,lastname,pay):
        self.firstname = firstname
        self.firstname = lastname
        self.pay = pay
        self.email = firstname+"."+lastname+"@gmail.com"
    
    def print_full_name(self):
        print(self.firstname+" "+self.lastname)


emp = Employee("shrinbas","mahanta","2000000")


emp.print_full_name()
        
    