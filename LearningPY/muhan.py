num: int
for num in range(1,(1+(int(input("Enter the checking limit:"))))):
    dup=num
    temp=0

    while num!=0:
        temp=((num%10)**(num%10)+temp)
        num=num//10
    if temp==dup:
        print("The number",temp,"is a muhan's number")