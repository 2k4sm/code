class Apollo:
    #veriable defination
    dest = "moon"
    #member functions
    def fly(self):
        print("spaceship flying....")
        
    def get_dest(self):
        print(f"Destination is: {self.dest}")


# objects one
obj_first = Apollo()
obj_second = Apollo()

#changing the destination of object one
obj_first.dest = "mars"
#calling of fly function
obj_first.fly()
#calling of get_dest func
obj_first.get_dest()
obj_second.fly()
obj_second.get_dest()