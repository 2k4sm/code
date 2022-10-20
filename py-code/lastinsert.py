list = []
n = int(input("Enter the number of integers to insert:"))

for i in range(0,n):
    num1=int(input("Enter the element to insert:"))
    list.append(num1)
print("The list formed is",list)

insert=int(input("Enter the integer to insert:"))
list.append(insert)

print("The new list is:",list)
