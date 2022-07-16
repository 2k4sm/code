def pcheck(x):
    y = int(x**0.5)
    for i in range(2,y):
        if x%i !=0 :
            return "prime"
        else:
            return "not_prime"
print(pcheck(7))