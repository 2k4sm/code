package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	//fmt.Println(arg())
	var (
		plaintext string
	)
	//checks the cmd line argument.
	arg()
	//takes the plaintext as user input using bufio module.
	fmt.Print("plaintext:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	plaintext = scanner.Text()
	//fmt.Println(plaintext)
	//Prints the ciphered text to the user as output.
	fmt.Printf("ciphertext:%s\n", cipher(plaintext))
	os.Exit(0)
}

//creating a function to cipher using the argument given in cmd line.
func cipher(plaintext string) string {
	var (
		newcip  string
		ind     int
		lowtext string
	)
	alpha := "abcdefghijklmnopqrstuvwxyz"
	cip := strings.ToLower(arg())
	//converts the plaintext to lower for comparison.
	lowtext = strings.ToLower(plaintext)
	for j, i := range lowtext {
		ind = strings.Index(alpha, string(i))
		//error handling for unexpected characters.
		if ind != -1 {
			//converting the case of the characters to as they were before changing.
			if unicode.IsUpper(rune(plaintext[j])) {
				newcip += strings.ToUpper(string(cip[ind]))
			} else {
				newcip += string(cip[ind])
			}
		} else {
			newcip += string(i)
			continue
		}

	}
	//returning the ciphered text.
	return newcip
}

//function to take the cmd line argument..
func arg() string {
	var (
		arg1 []string
		arg2 string = "Usage: ./substitution Key"
	)
	arg1 = os.Args[1:]
	//checks to take only one cmd line argument.
	for j := range arg1 {

		if j == 1 {
			fmt.Println("Usage: ./substitution Key")
			os.Exit(1)

		} else {
			arg2 = arg1[0]
			//checks if the key is of 26 letters or not.
			if len(arg2) != 26 {
				fmt.Println("Key must contain 26 letters.")
				os.Exit(2)
			} else {
				//checks if the ciphering key is all letter or not.
				for _, i := range arg2 {
					if unicode.IsLetter(i) == false {
						fmt.Println("Usage: ./substitution key")
						os.Exit(3)
					} else {
						//if everything goes well the key is stored
						arg2 = arg1[0]
					}

				}

			}
		}
	}
	//the cipher key is returned to the main function as a userinput.
	return arg2
}
