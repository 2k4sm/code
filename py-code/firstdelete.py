list = []
n = int(input("Enter the number of integers to insert:"))

for i in range(0,n):
    num1=int(input("Enter the element to insert:"))
    list.append(num1)
print("The list formed is",list)

num2 =0
list.append(num2)
temp=0

#print(list)
list[0],list[len(list)-1] = list[len(list)-1],list[0]
#print(list)
list.pop()
#print(list)

for j in range(1,len(list)+1):
    list[0],list[-j] = list[-j],list[0]
print(list)
list.pop()
print("The list after deleting the first letter is:",list)
