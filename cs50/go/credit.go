package main

import (
	"fmt"
	"math"
)

/*
	This is the implementation of luhn's algorithm used to verify the validity of credit/debit cards.

The key aspects for implementation:

	*starting from second to last number all the alternative numbers will be multiplied by 2 and added.(we will get sum1 here.)
	*Then all the left out numbers will give us the another addition.(sum2)
	*If the addition of these two numbers (sum1+sum2) give a number having 0 in it's unit place...then the card is a having a valid card number.
*/
func luhnal(card int64) int64 {
	var (
		chksum2 int64 = 0
		chksum1 int64 = 0
	)
	//This is the Luhn's algorithm.
	for card > 0 {
		chksum1 += card % 10
		card = card / 10
		scard := 2 * (card % 10)
		if scard > 9 {
			for scard > 0 {
				chksum2 += scard % 10
				scard = scard / 10

			}
			card = card / 10
		} else {
			chksum2 += scard
			card = card / 10
		}

	}
	return chksum1 + chksum2
}

// These are some useless codes to.....print  these statements...
func evalcard(lunnum int64) {
	if lunnum%10 == 0 {
		fmt.Println("The card's number satisfies the checksum....")

	} else {
		fmt.Println("The card is having a invalid number....")
	}

}

// This determines what type of card it is....
func whichcard(card int64) {
	mascar := card / int64(math.Pow10(14))
	amex := card / int64(math.Pow10(13))
	visa1 := card / int64(math.Pow10(12))
	visa2 := card / int64(math.Pow10(15))
	if amex == 34 || amex == 37 {
		fmt.Printf("The card with number %d is a AMEX card.\n", card)
	} else if mascar == 51 || mascar == 52 || mascar == 53 || mascar == 54 || mascar == 55 {
		fmt.Printf("The card with number %d is a MasterCard Card.\n", card)
	} else if visa1 == 4 {
		fmt.Printf("The card with number %d is a Visa card.\n", card)
	} else if visa2 == 4 {
		fmt.Printf("The card with number %d is a Visa card.\n", card)
	} else {
		fmt.Printf("The card with number %d is Invalid card.\n", card)
	}
}
func main() {
	var (
		card int64
	)
	//Taking the card number as user input.
	fmt.Printf("Enter the card number:")
	fmt.Scanf("%d", &card)
	for card < 1 {
		fmt.Printf("Enter the card number(+ve):")
		fmt.Scanf("%d", &card)
	}
	//some useless line of code because.....why not!!
	evalcard(luhnal(card))
	//when to check for card..
	if luhnal(card)%10 == 0 {
		whichcard(card)
	}

	//fmt.Println(luhnal(card))

}
