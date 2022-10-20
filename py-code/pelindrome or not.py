num1 = int(input("Enter your integer to check:"))
temp = 0
rev = 0
dup = num1
while num1 > 0:
    temp = num1 % 10
    rev = (rev * 10) + temp
    num1 = num1 // 10
if rev == dup:
    print("The entered integer is a palindrome.")
else:
    print("The entered integer is not a palindrome")
print(num1//10)
print(rev)