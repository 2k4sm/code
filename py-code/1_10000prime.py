count=0
for i in range(2,100001):
    p = 2
    num2 = 0
    while p < i:
        num2 = i % p
        p += 1
        if num2 == 0:
            break
    if num2 != 0:
        print((i))
    count+=1
print(count)
