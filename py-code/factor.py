import array

num1=int(input("Enter the number whose prime factor is to be found:"))

i=1
z=0
while i<=num1:
    if num1%i==0:
        print(i)
    i+=1