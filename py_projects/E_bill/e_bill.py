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
        self.tax = 0
        self.counter = 0
    
    def stripwatt(self):
        self.wattage = []
        watt = self.watt
        while watt != 0:
            if watt >= 10:
                for i in range(0,(watt//10)):
                    self.wattage.append("10kw")
                watt = watt%10
            elif watt >= 5:
                for i in range(0,(watt//5)):
                    self.wattage.append("5kw")
                watt = watt%5

            elif watt >= 3:
                for i in range(0,(self.watt//3)):
                    self.wattage.append("3kw")
                watt = watt%3

            elif watt >= 1:
                for i in range(0,(watt//1)):
                    self.wattage.append("1kw")
                watt = watt%1
            self.counter +=1


        

    def measure(self,gst):
        if self.counter == 0:
            Bill.stripwatt(self)
            self.tax = (gst/100)*self.watt
            for i in self.wattage:
                self.total += Bill.price[i]
            
            self.total += self.tax

            print(f'Tax:{self.tax}')
            print(f'Wattage:{self.wattage}')
            print(f'Total:{self.total}')
        else:
           print(f"Already measured TOTAL:{self.total}")


            

