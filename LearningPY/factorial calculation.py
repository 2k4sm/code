#print the factorial of a number.
a=int(input("Enter your number:"))
x=1
if a<0:
    print("It does not contain any factorial")
elif a==0:
    print("The factorial is 1")
else:
    if a>0:
        while (a>0):
            x=x*a
            a=a-1
    print("The factorial is",x)
