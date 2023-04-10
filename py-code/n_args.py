#Take a list and return a arranged list...

tlen=int(input("Enter the total numbers to be arranged:"))

list=[]

for i in range(0,tlen):
    list.append(int(input("Enter the elements present in the unarranged list:")))

#print("Your list is:",list)

temp=0

while tlen!=0:
    for j in range(0,len(list)-1):
        if list[j]>list[j+1]:
            list[j],list[j+1]=list[j+1],list[j]
    tlen-=1
print("The arranged list is:",list)
