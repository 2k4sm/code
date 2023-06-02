a=int(input("Enter the number:"))

if a==2:
    print("The number 2 is prime")

for i in range(2,a):
    if a%i==0:
        print("The number",a,"is a composite number:")
        continue
    else:
        print("The number",a,"is prime")
        break
