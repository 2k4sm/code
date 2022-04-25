list=[]
m=int(input("specify how many numbers to insert in your list:"))
for i in range(0,m):
    numl=int(input("Enter the list element:"))
    list.append(numl)
print("Your list is:",list)

for j in range(0,len(list)):
    if list[j]%5==0 and list[j]%7!=0:
        print("The list element",list[j],"is divisible by 5 but not divisible by 7.")
