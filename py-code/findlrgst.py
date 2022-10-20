dict = {1: int(input("Enter the first number:")),
        2: int(input("Enter the second number:")),
        3: int(input("Enter the last number:"))}
temp = 0

for i in range(1, 4):
    if dict[i] > temp:
        temp = dict[i]
print("The largest number among them is:", temp)
