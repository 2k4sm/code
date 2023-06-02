
import timeit

class sorter:
    def __init__(self,list):
        self.list = list  
        self.swap = 0  
    def bubble(self):
        for i in range(0,len(self.list)-1):
            for i in range(0,len(self.list)-1):
                if self.list[i] > self.list[i+1]:
                    self.list[i],self.list[i+1] = self.list[i+1],self.list[i]
                else:
                    continue
        return self.list
    def selection(self):
        for i in range(len(self.list)):
            min = i
            for j in  range(i+1,len(self.list)):
                if self.list[j] < self.list[min]:
                    min = j
                    self.swap = 1
            if self.swap == 1 :
                self.list[min],self.list[i] = self.list[i],self.list[min]
        return self.list
                    
