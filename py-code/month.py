mnum = int(input("Enter the month number to check for(1-12):"))
if mnum ==0 or  mnum > 12:
    print("You have entered an out of range input"
          " consider rectifying the mistake while re running the program")
mondict = { 1: "january,31days",
            2: "february,28days",
            3: "March,31days",
            4: "april,30days",
            5: "May,31days",
            6: "June,30days",
            7: "July,31days",
            8: "August,31days",
            9: "September,30days",
            10: "October,31days",
            11: "November,30days",
            12: "December,31days"}

for i in range(1,13):
    if i == mnum:
        out=mondict[i]
        print("Your desired month and it's days are:",out)
