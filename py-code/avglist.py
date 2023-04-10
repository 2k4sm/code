list=[]
m=int(input("specify how many numbers to insert in your list:"))
for i in range(0,m):
    numl=int(input("Enter the list element:"))
    list.append(numl)
print("Your list is:",list)
temp=0
for j in range(0,len(list)):
    temp+=list[j]
#print("sum of all digits is:",temp)

avg=temp/len(list)
print("The average of all the elements in the given list is:",avg)
