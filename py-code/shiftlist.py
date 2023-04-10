list=[]
m=int(input("specify how many number's to insert:"))
for i in range(0,m):
    numl=int(input("Enter the list element:"))
    list.append(numl)
print("Your list is:",list)

n=int(input("Enter the how many number's to shift:"))
for j in range(0,n):
    for k in range(0,len(list)-1):
        temp=0
        temp=list[k]
        list[k]=list[k+1]
        list[k+1]=temp
print("After shifting your list is:",list)
