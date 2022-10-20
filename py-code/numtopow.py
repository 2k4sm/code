
num1=int(input("Enter your integer:"))
dup=num1
dup2=num1
temp=0
count=0
while num1!=0:
    num1=num1//10
    count+=1

while count!=0:
    temp=((dup2%10)**count)+temp
    count=count-1
    dup2=dup2//10
if temp==dup:
    print("success",temp)
else:
    print("fail",temp)

