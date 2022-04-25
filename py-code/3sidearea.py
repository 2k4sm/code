l1=int(input("length of the first side of the triangle:"))
l2=int(input("length of the second side of the triangle:"))
l3=int(input("length of the third side:"))
s=(l1+l2+l3)/2

area= (s*(s-l1)*(s-l2)*(s-l3))**0.5

print("Area of the triangle is", area)
