num1=0
while num1<10001:
    dup=num1
    dup2=num1
    dup3=num1
    temp=0
    count=0
    while dup3!=0:
        dup3=dup3//10
        count+=1
    while count!=0:
        temp=((dup2%10)**count)+temp
        count=count-1
        dup2=dup2//10
    if temp==dup:
        print("success",temp)
    num1+=1