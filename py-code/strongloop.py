for num1 in range(1,(1+(int(input("Enter the checking limit:"))))):
    temp = 0
    l = 0
    num=num1
    dup=num
    dup2=num

    while dup2 != 0:
        dup2 = dup2 // 10
        l += 1

    while num != 0:
        temp = (((num % 10) ** l) + temp)
        num = num // 10
    if temp==dup:
        print("The number",temp,"is an armstrong number")

