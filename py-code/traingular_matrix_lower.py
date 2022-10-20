# make upper triangular matrix.
# make lower triangular matrix.
# transpose of a matrix.
#find trace of the matrix.

# THIS IS A LOWER TRIANGULAR MATRIX
list1 = []

for i in range(0,3):
    listg =[]
    for j in range(0,3):
        listg.append(int(input("Enter element:")))
        continue
    for i in range(0,1):
        list1.append(listg)
#print(list1)

for i in list1:
    for j in i:
        if i.index(j)>list1.index(i):
            print(0, end = " ")
        else:
            print(j, end = " ")
    print()
