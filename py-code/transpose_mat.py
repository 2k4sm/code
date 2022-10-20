import transpose
import matmkr


n = int(input("length of the matrix rows:"))
list1 = []
for i in range(0,n):
    list1.append(list(eval(input(f"Enter the {1} th list element:"))))



transpose = transpose.transpose(list1)
print(transpose)
