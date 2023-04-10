num1 =int(input("Enter your number to reverse:"))
temp = 0
rev = 0
dup = num1
while num1 > 0:
    temp = num1 % 10
    rev = (rev * 10) + temp
    num1 = num1 // 10
print(rev)
