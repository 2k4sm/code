list=[10,20,30,40]
p=int(input("Enter the position at which the number is to be inserted:"))
list+=[100]
p-=1
#print(len(list))
for i in range(p,len(list)):
    temp=0
    temp=list[i]
    list[i]=list[len(list)-1]
    list[len(list)-1]=temp
print(list)
