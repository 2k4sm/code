num1=1
num2=0
print(num2)
print(num1)

sum=num1+num2
print(sum)

while sum<10000:
    num2=num1
    num1=sum
    sum=num1+num2
    print(sum)