num1=int(input("Enter the value for the hypotenuse:"))
num2=int(input("Enter the value for the height:"))
num3=0

num3=(num1**2 - num2**2)**0.5

if (int(num3))==(float(num3)):
    print("The base of the right triangle is", num3)
else:
    print("The height of this config doesn't exist.")

