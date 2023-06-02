#Take three integers as input and print the output.
a=int(input("Enter first number:"))
b=int(input("Enter second number:"))
c=int(input("Enter third number:"))

if a>b>c:
    print("'1st no.' is the greatest number among all.")
elif a>c>b:
     print("'1st no.' is the greatest number among all.")
elif b>a>c:
    print("'2nd no.' is the greatest number among all.")
elif b>c>a:
     print("'2nd no.' is the greatest number among all.")
elif c>b>a:
    print("'3rd no. ' is the greatest number among all.")
elif c>a>b:
     print("'3rd no.' is the greatest number among all.")
else:
    print(" all the numbers are equal.")
    