# checks if a year is a leap year or not.


num1 = int(input("Enter the year to check for:"))

if num1 % 4 == 0:
    if num1%100!=0:
        print("The year",num1,"is a leap year.")
    elif num1%100==0 and num1%400==0:
        print("The year",num1,"is a leap year.")
    else:
        print("The year",num1,"is not a leap year.")
else:
    print("The year",num1,"is not a leap year.")
