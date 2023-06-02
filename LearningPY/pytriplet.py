i=0
j=0
k=0

for i in range(1,1000):
    i+=1
    for j in range(1,i):
        j+=1
        for k in range(1,j):
            if (k**2+j**2)==i**2:
                print(i,j,k,"are pythagorean triplets.")
            k+=1