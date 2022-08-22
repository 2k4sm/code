package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	//user defined function to take command line arguments.
	arg()
	//To take string as user input.
	fmt.Print("plaintext:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	plaintext := scanner.Text()

	//To display the ciphered text.
	fmt.Printf("ciphertext:%s\n", cipher(plaintext))
	os.Exit(0)
}

//defination of the function for taking cmd-line arguments and convert them to integer..also to check for valid command line inputs. or terminate
//the programme otherwise with exit code (1)

func arg() int {
	var (
		intarg int
	)
	arg1 := os.Args[1:]
	for a, i := range arg1 {
		if a != 0 {
			fmt.Println("Usage: ./ceaser key")
			os.Exit(1)
		} else {
			intarg, _ = strconv.Atoi(i)
			for intarg == 0 {
				fmt.Printf("Usage: ./ceaser key\n")
				os.Exit(1)
			}
		}

	}
	return intarg

}

//This is the implementation of ceaser's cipher here the formula for ciphering is: c = ((i+k)%26)
func cipher(plaintext string) (cip string) {
	duptext := plaintext
	plaintext = strings.ToLower(plaintext)
	var cip1 rune
	cip = ""
	for _, i := range plaintext {
		cip1 = rune((((int(i) + arg()) - 97) % 26) + 97)
		cip += string(cip1)
	}
	var cip2 string
	for a, i := range duptext {
		if unicode.IsUpper(i) {
			newup := string(cip[a])
			cip2 += strings.ToUpper(newup)
		} else {
			cip2 += string(cip[a])
		}
	}
	return cip2
}
