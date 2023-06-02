num1=int(input("Enter your first number:"))
num2=int(input("Enter your second number:"))

#for hcf:

temp=1
hcf=1
if num1>=num2:
    temp=num2
elif num2>=num1:
    temp=num1

while temp>0:
    if num1%temp==0 and num2%temp==0:
        hcf=temp
        break
    temp-=1
print("The hcf is",hcf)
#for lcm:

lcm=(num1*num2)/hcf
print("The lcm is:",lcm)

