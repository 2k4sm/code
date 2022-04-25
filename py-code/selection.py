list=[]
n=int(input("Enter the number of elements to add:"))
for i in range(0,n):
    numin=int(input("Enter the numbers to input:"))
    list.append(numin)
print("The list before sorting:",list)

n=0
for i in range(0,len(list)):
    for j in range(0,i+1):
        if list[i]>list[j]:
            temp=0
            temp=list[j]
            list[j]=list[i]
            list[i]=temp
    n+=1
    #print("After procesing", n, "Step", "The list is:", list)
print("The list after sorting:",list)