dict={1:int(input("Enter the present day(dd):")),
      2:int(input("Enter the present month(mm):")),
      3:int(input("Enter the present year(yyyy):")),

      4:int(input("Enter the birth day(dd):")),
      5:int(input("Enter the birth month(mm):")),
      6:int(input("Enter the birth year(yyyy):")),
      }

yrd=(dict[3]-1) -(dict[6])
lpy=yrd//4
yrd=yrd*365+lpy
print(yrd)
lyr=dict[3]//100
iyr=dict[6]//100
count=0
for i in range(iyr+1,lyr+1):
    if i%4!=0:
        count=count+1
yrd=yrd-count

print(yrd)
print(count)




