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

			fmt.Println("Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet. Nisi anim cupidatat excepteur officia. Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate voluptate dolor minim nulla est proident. Nostrud officia pariatur ut officia. Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem duis laboris cupidatat officia voluptate. Culpa proident adipisicing id nulla nisi laboris ex in Lorem sunt duis officia eiusmod. Aliqua reprehenderit commodo ex non excepteur duis sunt velit enim. Voluptate laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa duis.)
			if guess < tguess {
				fmt.Printf("The guess is smaller by %d \n", tguess-guess)

			} else {
				fmt.Printf("The guess is greater by %d \n", guess-tguess)
			}

		} else {
			fmt.Printf("you guessed it in %d tries..", count)
			fmt.Printf("You guessed correct that is %d \n", guess)
			break
			//This is a change in the file.
		}

	}

}

