#for prime number:
# take a number
# run a loop upto that number dividing each number. if divisible by any other number other than one and itself.
# Then not prime.

num1=int(input("Enter the number to check for:"))
if num1==2 or num1==1:
    print("The number is prime.")
else:
    p=2
    num2=0
    while p<num1:
        num2=num1%p
        p+=1
        if num2==0:
            break
    if num2!= 0:
        print("The number is a prime number.")
    else:
        print("The number is not prime")


