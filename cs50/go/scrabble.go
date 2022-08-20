/* Primary objective is to implement three things in the program...
*The first is to take user input as a string...
*Then score each letter of the input from the score array defined and add the scores...
*Then compare the scores to see who is the winner...
 */

package main

import (
        "fmt"
        "unicode"
        "strings"

        )

func main() {
    var(
        usrinp1 string
        usrinp2 string
        score1 int
        score2 int
    )
    //Taking userinput from player 1
    fmt.Printf("Player 1 Enter your word:")
    fmt.Scanf("%s",&usrinp1)
    usrinp1 = Isupper(usrinp1)
    //Taking userinput from player 2
    fmt.Printf("Player 2 Enter your word:")
    fmt.Scanf("%s",&usrinp2)
    usrinp2 = Isupper(usrinp2)
    //Calculating the score
    score1 =compute(usrinp1)
    score2 = compute(usrinp2)
    //Checking who wins
    if score1 > score2 {
        fmt.Printf("Player 1 Wins!!:Score = P1:%d,P2:%d\n",score1,score2)
    }else if score2 > score1 {
        fmt.Printf("Player 2 Wins!!:Score = P1:%d,P2:%d\n",score1,score2)
    }else {
        fmt.Println("Tie!!")
    }

    //fmt.Println(score1,score2)
    //fmt.Println(usrinp1,usrinp2)

}
//This uses the unicode.isupper and strings.tolower to convert uppercase inputs to lowercase..
func Isupper(str2 string) (string){
    for _,i := range str2 {
        if unicode.IsUpper(i) {
            str2 = strings.ToLower(str2)
            break
        }
    }
    return str2
}
//This uses comparison between characters to determine the character and it's points and return the point in the sum..
//After that it do it  with all the characters from the user input and adds the point to the sum..
func compute(str1 string) (int) {
    Points := [26]int{1,3,3,2,1,4,2,4,1,8,5,1,3,1,1,3,10,1,1,1,1,4,4,8,4,10}
    chars  := "abcdefghijklmnopqrstuvwxyz"
    sum := 0
    for _,i := range str1 {
        for _,j := range chars {
            if i == j {
                index := strings.Index(chars, string(i))
                sum += Points[index]
            }
        }
    }
    return sum
}
/*
BUG
the compute is failing when userinputs are having repeating characters in them.
*/
