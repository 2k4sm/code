from sorting import *
class searcher:

    def __init__(self,list1,item):
        self.item = item
        self.list = list1
        self.first = 0
        self.last = len(self.list)
        self.mid = (self.first+self.last)//2
    def binsrh(self):
        for i in range(self.last):
            if self.item < self.list[self.mid]:
                self.last = self.mid-1
            elif self.item > self.list[self.mid]:
                self.first = self.mid+1
            elif self.item == self.list[self.mid]:
                print(f"The position of the {self.item} is {self.mid}.")
            else:
                print("not found")


number1 = searcher(sorter([2,3,4,5,6,7,8,9]).selection(),5)
number1.binsrh()
