# comparing n numbers to find the greatest...
tlen=int(input("Enter the total numbers to compare:"))

list=[]

for i in range(0,tlen):
    list.append(int(input("Enter your number to add:")))

#print("Your list is:",list)

temp=0
for i in range(0,len(list)):
    if list[i]>temp:
        temp=list[i]
print("The greatest number is:",temp)
