def search(list1,item):
    low = 0
    high = len(list1) - 1
    
    while low<=high:
        mid  = (low+high)/2
        guess = list1[mid]

        if guess == item:
            return mid
        if guess > item:
            high = mid - 1
        else:
            low = mid+1
    return None


listlen = int(input("Enter List len:"))
i=0
while i < listlen:
    list1 = []
    while True:
        if i == listlen:
            break
        else:
            i+=1
        input1 = int(input("List Element:"))
        list1.append(input1)
        

item = int(input("Enter the number to search for:"))

search(list1,item)

        