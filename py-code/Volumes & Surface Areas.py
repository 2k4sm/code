# VOLUMES & SURFACE AREAS OF 3D SHAPES
e=1
while e!=0:
  print("TO CALCULATE VOLUMES & SURFACE AREAS OF FOLLOWING SHAPES")
  print("1. Cube")
  print("2. Cuboid")
  print("3. Cylinder")
  print("4. Cone")
  print("5. Sphere")
  print("6. Hemisphere")
  print("0. Exit")

  c = int(input("Type the Seriel No of the shape: "))
  if c>0 and c<7:
    if c==1:
      shape="CUBE"
      a=int(input("Enter the value of SIDE of the CUBE: "))
      vol=a**3
      sa=6*a**2
    elif c == 2:
      shape="CUBOID"
      l=int(input("Enter the value of LENGTH of the CUBOID: "))
      w=int(input("Enter the value of WIDTH  of the CUBOID: "))
      h=int(input("Enter the value of HEIGHT of the CUBOID: "))
      vol=l*w*h
      sa=l*w+l*h+w*h
      sa=2*sa
    elif c == 3:
      shape="CYLINDER"
      r=int(input("Enter the value of RADIUS of the CYLINDER: "))
      h=int(input("Enter the value of HEIGHT of the CYLINDER: "))
      vol=22/7*r**2*h
      sa=r+h
      sa=2*22/7*r*sa
    elif c == 4:
      shape="CONE"
      r=int(input("Enter the value of RADIUS of the CONE      : "))
      h=int(input("Enter the value of HEIGHT of the CONE      : "))
      vol=1/3*22/7*r**2*h
      s1=(r**2+h**2)**0.5
      sa=r+s1
      sa=22/7*r*sa
    elif c ==5:
      shape="SPHERE"
      r=int(input("Enter the value of RADIUS of the SPHERE: "))
      vol=4/3*22/7*r**3
      sa=4*22/7*r**2
    elif c ==6:
      shape="HEMISPHERE"
      r=int(input("Enter the value of RADIUS of the HEMISPHERE: "))
      vol=2/3*22/7*r**3
      sa=3*22/7*r**2
    print ("Choosen shape is ",shape)
    print ("Volume of        ",shape,"is =",vol)
    print ("Surface Area of  ",shape,"is =",sa)
    print ("")
    str =input("Press ENTER KEY to continue")
  elif c>6:
    print("")
    print("Please enter valid serial number of the SHAPE")
    print("")
  e=c
  if e==0:
    break
  
    

