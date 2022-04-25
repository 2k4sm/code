# program to take two integers a operator and compute them.
a=int(input("Enter first number:"))
b=int(input("Enter second number:"))

print("ADD=1")
print("SUB=2")
print("MULTI=3")
print("DIVI=4")
print("choose from above according to what you want to do:")

add=a+b
sub=a-b
multi=a*b
divi=a/b

c=int(input("what to do:"))

if c==1:
    print("your result is",add)
elif c==2:
    print("your result is",sub)
elif c==3:
    print("your result is",multi)
elif c==4:
    print("your result is",divi)
else:
    print("operator unspecified")




