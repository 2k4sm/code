num1=int(input("Enter your number whose sum has to be found:"))
temp=0
while num1!=0:
    temp=(num1%10)+temp
    num1=num1//10
print(temp)
