divdnd=int(input("Enter your dividend:"))
divisr=int(input("Enter your divisor:"))
remainder=1
while divisr>0:
    remainder=divdnd%divisr
    if remainder==0:

        hcf=divisr
        break
    else:
        dividnd=divisr
        divisr=remainder
print(hcf)
