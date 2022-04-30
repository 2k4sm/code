// program prints if the number is positive,negative or zero.

package main

import "fmt"

func main(){
  var num1 int64

  fmt.Println("Enter your number:")
  fmt.Scan(&num1)

  if num1>0{
    fmt.Printf("The given Number %v is Positive.",num1)
  }else if num1==0{
    fmt.Printf("The number entered is %v",0)
  }else {
    fmt.Printf("The number %v is negative.",num1)
  }

}
