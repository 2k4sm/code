list = []
n = int(input("Enter the number of integers to insert:"))

for i in range(0,n):
    num1=int(input("Enter the element to insert:"))
    list.append(num1)
print("The list formed is",list)

#num1 = int(input("Enter the integer to insert in the first place is:"))
list.append(num1)
temp=0
for i in range(0,len(list)):
    temp = list[i]
    list[i] = list[len(list)-1]
    list[len(list)-1] = temp
#print("The new inserted list is:",list)
