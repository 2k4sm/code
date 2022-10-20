word = str(input("Enter your string to create the pattern:"))
t = ()
for i in range(0,len(word)):
    c1 = 0
    while c1!=3:
        req = word[0+i:len(word)]
        tup = (req,)
        t+=tup
        c1+=1

print(t)

