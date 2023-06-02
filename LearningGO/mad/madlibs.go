package main

import "fmt"

func main(){
  var adjective string
  var noun string

  fmt.Println("Enter the adjective:")
  fmt.Scanf("%s", &adjective)

  fmt.Println("Enter the noun:")
  fmt.Scanf("%s", &noun)

  fmt.Printf("It was a %s cold november day",adjective)
  fmt.Printf("I woke up to the smell of %s",noun)


}
