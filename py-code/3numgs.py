# find the greatest between three numbers...

num1=int(input("Enter your first number:"))
num2=int(input("Enter your second number:"))
num3=int(input("Enter your third number:"))

temp=0


if num1>num2:
    temp=num1
else:
    temp=num2

if num3>temp:
    print("The greatest number is:",num3)
else:
    print("The greatest number is:",temp)
    
