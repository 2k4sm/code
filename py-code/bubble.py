list=[]
n=int(input("Enter the number of elements to add:"))
for i in range(0,n):
    numin=int(input("Enter the numbers to input:"))
    list.append(numin)
print("The list before sorting:", list)

n=0
while n<(len(list)-1):
    for i in range(0,len(list)-1):
        if list[i]>list[i+1]:
            temp=0
            temp=list[i]
            list[i]=list[i+1]
            list[i+1]=temp
    #print("After procesing",n+1,"Step","The list is:",list)
    n+=1
print("The list after sorting is:",list)