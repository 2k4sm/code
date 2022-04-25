num1=1
num2=0
lower=int(input("Enter your lower:"))
upper=int(input("Enter your upper:"))
sum=num1+num2
flag=1

while sum<10000:
    num2=num1
    num1=sum
    sum=num1+num2
    if sum>lower and sum<upper :
        print(sum)
        flag=0

if flag!=0:
    print("No elements are present between", lower,"and",upper)