#check if three numbers are triplet or not

num1=int(input("Enter the first number to check for triplet:"))
num2=int(input("Enter the second number to check for  triplet:"))
num3=int(input("Enter the third number to check for  triplet:"))
if num1<num2 and num1<num3:
    if num2<num3 and (num1+num2)>num3:
        if (num1**2 +num2**2)==num3**2:
            print(num1,num2,num3,"are pythagorean triplets.")
        else:
            print("The provided numbers are not triplets.")
    elif num3<num2 and (num1+num3)>num2:
        if  (num1**2 +num3**2)==num2**2:
            print(num1,num2,num3,"are pythagorean triplets.")
        else:
            print("The provided numbers are not triplets.")

if num2<num1 and num2<num3:
    if num1<num3 and (num1+num2)>num3:
        if (num1**2 +num2**2)==num3**2:
            print(num1,num2,num3,"are pythagorean triplets.")
        else:
            print("The provided numbers are not triplets.")
    elif num3<num1 and (num3+num2)>num1:
        if  (num3**2 +num2**2)==num1**2:
            print(num1,num2,num3,"are pythagorean triplets.")
        else:
            print("The provided numbers are not triplets.")

if num3<num1 and num3<num2:
    if num1<num2 and (num1+num3)>num2:
        if (num1**2 +num3**2)==num2**2:
            print(num1,num2,num3,"are pythagorean triplets.")
        else:
            print("The provided numbers are not triplets.")
    elif num2<num1 and (num2+num3)>num1:
        if  (num3**2 +num2**2)==num1**2:
            print(num1,num2,num3,"are pythagorean triplets.")
        else:
            print("The provided numbers are not triplets.")