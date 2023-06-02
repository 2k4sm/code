list=[]
m=int(input("specify how many numbers to insert in your list:"))
for i in range(0,m):
    numl=int(input("Enter the list element:"))
    list.append(numl)
print("Your list is:",list)
temp=list[0]
for j in range(0,len(list)):
    if temp>list[j]:
        temp=list[j]
print("The smallest number in the list is:",temp)
