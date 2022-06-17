package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var tguess int64
	fmt.Printf("Enter the secret number!!!(keep it btwn 1-100):")
	fmt.Scanf("%d", &tguess)
	
	var guess int64
	var count int = 0
	
	for {
		count += 1
		guess = rand.Int63n(101)
		if guess != tguess {

			fmt.Printf("The guess is %d which is wrong!!! \n", guess)
			if guess < tguess {
				fmt.Printf("The guess is smaller by %d \n", tguess-guess)

			} else {
				fmt.Printf("The guess is greater by %d \n", guess-tguess)
			}

		} else {
			fmt.Printf("you guessed it in %d tries..", count)
			fmt.Printf("You guessed correct that is %d \n", guess)
			break
		}

	}

}
