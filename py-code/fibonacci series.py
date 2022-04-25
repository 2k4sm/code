a=int(input("Length of the series:"))
p,q=0,1
c=0
print("The fibonacci series is:")
while c<a:
    print(p)
    r=p+q
    p=q
    q=r
    c+=1


