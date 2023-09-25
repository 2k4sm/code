a=int(input("Enter first number:"))
b=int(input("Enter second number:"))
hcf=1
p=1


while (p<=a and p<=b):
    if (a%p==0 and b%p==0):
        hcf=p
    p+=1


print(hcf)


c=(a*b)/hcf
print("lcm is:",c)
