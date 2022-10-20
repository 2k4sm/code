// This is a implementation of a index which checks the grade level of a book.
// Coleman lieu index formula = index = 0.0588*l-0.296*s-15.8
// where l is avarage number of letters per 100 words and s is avarage number of sentences per word.
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
)

func main() {
	//User input string
	var (
		text string
	)

	fmt.Printf("Text:")
	//takes the text as user input.
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text = scanner.Text()
	//calculates the index/grade of the book
	index := cole_ind(text)
	//gives a round figure of the above index.
	round := math.Round(index)
	//prints the number of words,lettters and sentences.
	fmt.Printf(" Words:%.0f\n Letters:%.0f\n Sentences:%.0f\n", lett(text), word(text), sent(text))
	//checks if 1>round>13 and prints accordingly..
	if round > 13 {
		fmt.Println(" Grade:13+")

	} else if round < 1 {
		fmt.Println(" Grade:1-")
	} else {
		fmt.Printf(" Grade:%d\n", int(round))
	}
}

//this is a implementation of the coleman lieu index formula..
func cole_ind(text string) (index float64) {
	l := float64((lett(text) / word(text)) * 100)
	s := float64((sent(text) / word(text)) * 100)

	index = 0.0588*l - 0.296*s - 15.8
	return index
}

//calculates the number of letters.
func lett(text string) (lett float64) {
	lett = 0
	for _, i := range text {
		if is_alphanum(string(i)) {
			lett++
		}
	}
	return lett
}

//calculates the number of words.
func word(text string) (word float64) {
	word = 0
	for _, i := range text {
		if string(i) == " " {
			word++
		}
	}
	return word + 1
}

//calculates the number of sentences.
func sent(text string) (sent float64) {
	sent = 0
	for _, i := range text {
		if string(i) == "." || string(i) == "?" || string(i) == "!" {
			sent++
		}
	}
	return sent
}

//checks if a number is alphanumeric or not.
func is_alphanum(text string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(text)
}
