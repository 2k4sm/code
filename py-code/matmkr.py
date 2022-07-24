#Function defined to make a matrix

def matrix(new_lst):
    for i in new_lst:
        for j in i:
            print(j,end=" ")
        print()
    return new_lst

#DRIVER CODE


'''
list1 = [list(eval(input(f"Enter list elements for {1} th row:"))),
         list(eval(input(f"Enter list elements for {2} th row:"))),
         list(eval(input(f"Enter list elements for {3} th row:")))]


print(matrix(list1))
'''