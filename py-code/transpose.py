import matmkr


# function defined to create transpose of a matrix.

def intertranspose(new_lst):
    for i in new_lst:
        for j in i:
            ind1 = new_lst.index(i)
            ind2 = i.index(j)
            if ind1 >= ind2:
                new_lst[ind2][ind1],new_lst[ind1][ind2] = new_lst[ind1][ind2],new_lst[ind2][ind1]
    return new_lst


def transpose(new_lst):
    intertranspose(intertranspose(new_lst))
    return new_lst


'''
list1 = [list(eval(input(f"Enter the {1} th list element:"))),
         list(eval(input(f"Enter the {2} th list element:"))),
         list(eval(input(f"Enter the {3} th list element:")))]

print(transpose(list1))
matmkr.matrix(transpose(list1))
'''
