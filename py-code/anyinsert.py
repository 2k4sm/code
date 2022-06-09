list = []
n = int(input("Enter the number of integers to insert:"))

for i in range(0,n):
    num1=int(input("Enter the element to insert:"))
    list.append(num1)
print("The list formed is",list)

any1 = int(input("Enter your number to insert:"))
loc = int(input("Enter your desired location in the list to insert:"))
list.append(any1)
print(list)

for j in range(loc,len(list)):
    list[j-1],list[len(list)-1] = list[len(list)-1],list[j-1]
print(list)

