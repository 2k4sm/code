# to create an librarian for keeping track of money and days..
#charge for 2nd week up-10rs  5th-week up-25rs  7-week onwards--150rs
# fixed charge 200rs

num1=int(input("Enter the days for which the book has been taken for:"))
price=0
fixchr=250
week=num1//7

for i in range(1,week+1):
    if i<=2:
        price=fixchr
    if i<=5 and i>2:
        price=price+10
    if i<=7 and i>5:
        price=price+25
    if i>7:
        price=price+150


print("The amount you are required to pay is:",price)
