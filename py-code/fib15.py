num1=1
num2=0
print(num1)

sum=num1+num2
print(sum)
count=0
while count<15:
    num2=num1
    num1=sum
    sum=num1+num2
    print(sum)
    count=count+1