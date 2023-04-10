num=int(input("Enter the number whose binary is to be found:"))
dup1=num
n=0
temp=0
while dup1>0:
    rem=dup1%2
    dup1=dup1//2
    temp=temp+rem*(10**n)
    n+=1
print("The binary of the given number is:",temp)
