# write a program to take a username and a p code .ensure that p code doesn't contain username and matches 'Trident111'.

uname = input("Enter your username:")
pcode = (input("Enter code:"))

if uname not in pcode and pcode == 'Trident111':
    print("your code is valid to proceed.")

else:
    print("your code is invalid.")
