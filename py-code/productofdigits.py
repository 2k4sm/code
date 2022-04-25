num1=int(input("Enter your number whose product has to be found:"))
temp=1
while num1!=0:
    temp=(num1%10)*temp
    num1=num1//10
print(temp)
