
print("""
Please create a list to begin with...
======================================================================================================
""")


list = []
n = int(input("Enter the number of integers to insert in your list:"))

for i in range(0,n):
    num1=int(input("Enter the element to insert:"))
    list.append(num1)
print("The list formed is",list)


print("""
This is a program to do either insertion of an element or deletion of an element to or from the list..
======================================================================================================
1- Inserting an element
2- Deleting an element
""")
init = 0
init=int(input("Enter (1-2):"))
if init == 1:
    # There are three operations now
    # 1-inserting in the last place
    # 2-inserting in the first place
    # 3-inserting in the any place
    print("""
    Specify your insertion type:
    ===============================
    1-Inserting in the last place
    2-Inserting in the first place
    3-Inserting in any place
    ===============================
    """)
    init1 = int(input("Enter(1-3):"))
    if init1 == 1:
        insert=int(input("Enter the integer to insert:"))
        list.append(insert)

        print("The new list is:",list)
    elif init1 == 2:
        num1 = int(input("Enter the integer to insert in the first place is:"))
        list.append(num1)
        #print(list)

        temp=0
        for i in range(0,len(list)):
            temp = list[i]
            list[i] = list[len(list)-1]
            list[len(list)-1] = temp
        print("The new inserted list is:",list)
    elif init1 == 3:
        any1 = int(input("Enter your number to insert:"))
        loc = int(input("Enter your desired location in the list to insert:"))
        list.append(any1)
        #print(list)

        for j in range(loc,len(list)):
            list[j-1],list[len(list)-1] = list[len(list)-1],list[j-1]
        print("The new list formed is:",list)
    else:
        print("Invalid input given....")
elif init == 2:
    # There are three operations now
    # 1-Deleting in the last place
    # 2-Deleting in the first place
    # 3-Deleting in the any place
    print("""
    Specify your deletion type:
    ==============================
    1-Deletion in the last place
    2-Deletion in the first place
    3-Deletion in any place
    ==============================
    """)
    init1 = int(input("Enter(1-3):"))
    if init1 == 1:
        list.pop()
        print("The New list is",list)
    elif init1 == 2:
        num2 =0
        list.append(num2)
        temp=0

        #print(list)
        list[0],list[len(list)-1] = list[len(list)-1],list[0]
        #print(list)
        list.pop()
        #print(list)

        for j in range(1,len(list)+1):
            list[0],list[-j] = list[-j],list[0]
        #print(list)
        list.pop()
        print("The list after deleting the first letter is:",list)
    elif init1 == 3:
        rem = int(input("Enter the location of integer to remove:"))

        list.pop(rem-1)

        print("Your new list after deletion is:",list)
    else:
        ("Invalid input given...")
else:
    print("Invalid input given")
