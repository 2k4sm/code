def pel(num,temp = 0,count = 1):
    dup = num
    #rev = num
    while dup!=0:
        dup = dup//10
        count= (count*10)
        continue
    while num != 0:
        temp+=(num%10)*count
        num = num//10
        #rev//10
        count = count//10
        continue
    return "your Pelindrome is:",temp//10

num=int(input("Enter your number:"))
print(pel(num))

