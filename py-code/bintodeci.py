num=int(input("Enter your binary number:"))
temp=0
dup=num
n=0
#m=0
while dup!=0:
    temp=temp+dup%10*(2**n)
    dup=dup//10
    n+=1

print("The number for this binary representation is:",temp)
