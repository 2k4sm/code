# THIS IS A UPPER TRIANGULAR MATRIX
list1 = []

for i in range(0,3):
    listg =[]
    for j in range(0,3):
        listg.append(int(input(f"Enter{j+1}th element:")))
        continue
    for i in range(0,1):
        list1.append(listg)
#print(list1)

for i in list1:
    for j in i:
        if list1.index(i)>i.index(j):
            print(0, end = " ")
        else:
            print(j, end = " ")
    print()
