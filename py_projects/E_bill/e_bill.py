class Bill():

    price = {
        "1kw" : 5,
        "3kw" : 25,
        "5kw" : 50,
        "10kw" : 100
    }

    def __init__(self):
        self.watt = int(input("Watt:"))
        self.total = 0
        self.wattage = []
    
    def stripwatt(self):
        self.wattage = []
        while self.watt != 0:
            if self.watt >= 10:
                for i in range(0,(self.watt//10)):
                    self.wattage.append("10kw")
                self.watt = self.watt%10
            elif self.watt >= 5:
                for i in range(0,(self.watt//5)):
                    self.wattage.append("5kw")
                self.watt = self.watt%5

            elif self.watt >= 3:
                for i in range(0,(self.watt//3)):
                    self.wattage.append("3kw")
                self.watt = self.watt%3

            elif self.watt >= 1:
                for i in range(0,(self.watt//1)):
                    self.wattage.append("1kw")
                self.watt = self.watt%1
        

        def measure(self,gst,overdue):
            

