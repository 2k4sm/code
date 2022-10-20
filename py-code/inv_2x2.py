list1 = []

for i in range(0,2):
    listg =[]
    for j in range(0,2):
        listg.append(int(input("Enter:")))
        continue
    for i in range(0,1):
        list1.append(listg)

for r in list1:
    for j in r:
        print(j,  end ="  ")
    print()
print(list1)
#============================FINDING THE DETERMINANT==========================
for i in range(0,2):
    for j in range(0,1):
        if i == j:
            mul1 = list1[i][i]*list1[j+1][j+1]
        else:
            mul2 = list1[i][j]*list1[j][i]
z = mul1-mul2
print(f"The determinant of the given matrix is:{z}")
#============================FINDING THE ADJOINT MATRIX=======================
temp = 0
for i in range(0,2):
    for j in range(0,1):
        if i == j :
            temp = list1[i][j]
            list1[i][j] = list1[i+1][j+1]
            list1[i+1][j+1] = temp
            #print(list1)
        else:
            list1[i][j],list1[j][i] = -list1[i][j],-list1[j][i]
print(f"The adjoint of the given matirix is:{list1}")

#===========================FINDING INVERSE===================================
print("The inverse of the given matrix is:")
for r in list1:
    for j in r:
        print(j,"/",z,  end ="  ")
    print()
