a=int(input("Enter your number:"))
if a%5==0 and a%11==0:
    print("The number is divisible by 5 and 11")
elif a%5==0:
    print("The number is divisible by 5")
elif a%11==0:
    print("The number is divisible by 11")
else:
    print("The number is neither divisible by 5 nor divisible by 11")
