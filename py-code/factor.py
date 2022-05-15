num1=int(input("Enter the number whose factor is to be found:"))

i=1
while i<=num1:
    if num1%i==0:
        print(i)
    i=i+1
