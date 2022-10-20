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

for i in range(0,2):
    for j in range(0,1):
        if i == j:
            mul1 = list1[i][i]*list1[j+1][j+1]
        else:
            mul2 = list1[i][j]*list1[j][i]
print(f"The determinant of the given matrix is:{mul1-mul2}")