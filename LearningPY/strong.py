num =int(input("Enter your number:"))
dup = num
dup2=num
temp = 0
l = 0
while dup2 != 0:
    dup2 = dup2 // 10
    l += 1


while num != 0:
    temp = (((num % 10) ** l) + temp)
    num = num // 10

if dup==temp:
    print("The provided number is an armstrong number.")
else:
    print("The provided number is not an armstrong number.")
