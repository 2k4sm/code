# wap to create a 3*3 matrix..

list1 = []

for i in range(0,3):
    listg =[]
    for j in range(0,3):
        listg.append(int(input("Enter:")))
        continue
    for i in range(0,1):
        list1.append(listg)

for r in list1:
    for j in r:
        print(j,  end ="  ")
    print()
