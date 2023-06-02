def fact(a):
    temp=1
    b=(a-1)
    while a!=1:
        fac=a*b
        temp=(temp*fac)/b
        a-=1
    return 'The factorial is', temp

print(fact(int(input("Enter your integer for which factorial is to be found:"))))
