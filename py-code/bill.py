unit=int(input("Enter the units of electricity used:"))

if unit<=50:
    price=unit*1.2
elif unit>50 and unit<=150:
    price=50*1.2
    price=price+(unit-50)*1.75
elif unit>150 and unit<=250:
    price=50*1.2
    price=price+100*1.75
    price=price+(unit-150)*2.25
elif unit>250 and unit<=350:
    price=50*1.2
    price=price+100*1.75
    price=price+100*2.25
    price=price+(unit-250)*3.15
elif unit>350 and unit<=450:
    price = 50 * 1.2
    price = price + 100 * 1.75
    price = price + 100 * 2.25
    price = price + 100 * 3.15
    price = price +(unit-350)*4.15
else:
    price = 50 * 1.2
    price = price + 100 * 1.75
    price = price + 100 * 2.25
    price = price + 100 * 3.15
    price = price + 100 * 4.15
    price = price + (unit-450)*5

serv=(10/100)*price
print("service tax =",serv)
fuch=(20/100)*price
print("fuel charge =",fuch)


print("Your total price is",price+serv+fuch)





