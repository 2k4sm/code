str1=str(input("Enter the string:"))
for i in range(0,len(str1)):
    if str1[i]=="a" or str1[i]=="e" or str1[i]=="i" or str1[i]=="o" or str1[i]=="u" :
        str1=str1.replace(str1[i],"*")
print("The string after replacing the vowels with * is:",str1)
