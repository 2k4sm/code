
str1=str(input("Enter the string:"))
chr=str(input("Enter the character which we have to  check for:"))
count=0
for i in range(len(str1)):
    x=(str1[i])
    if x==chr:
        count=count+1
print("The number of times",chr,"appears in this given string is",count)
