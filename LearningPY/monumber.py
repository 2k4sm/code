import math

for num in range(1,10000000):
    dup = num
    rev=0
    while num != 0:
        temp=num%10
        rev=math.factorial(temp)+rev
        num=num//10
    if rev==dup:
        print("The",rev , "is a monumber.")
