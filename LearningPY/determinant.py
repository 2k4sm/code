def determinant(matrix):
    n = len(matrix)
    if(n == 1):
        return matrix[0][0]
    if(n == 2):
        return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
    det = 0
    for j in range(n):
        cofactor = matrix[0][j]*determinant([row[:j] + row[j+1:] for row in matrix[1:]])
        det += ((-1)**j)*cofactor
    return det


print(determinant([[2,0,0],[0,2,0],[0,0,2]]))