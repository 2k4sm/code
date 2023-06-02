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
#print(list1)
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
print(list1)

for r in list1:
    for j in r:
        print(j,  end ="  ")
    print()