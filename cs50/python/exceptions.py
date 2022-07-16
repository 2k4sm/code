import sys

try:
    x = int(input("Enter x:"))
    y = int(input("Enter y:"))

except ValueError:
    print("Error:Invalid input")
    sys.exit(1)
try:
    result = x/y
except ZeroDivisionError:
    print("Error:cannot divide by 0.")
    sys.exit(1)
print(f"The result of {x}/{y}={result}")
