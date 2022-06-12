package main

import "fmt"

func main(){
  fmt.Println("This is a program to check leap year for a given year...")
  var year int
  fmt.Printf("Enter the year:")
  fmt.Scanf("%d",&year)



  if year % 4 == 0 && year % 100 != 0 {
    fmt.Printf("The year %d is a leap year!\n",year)

  }else if year % 400 ==0{
    fmt.Printf("The year %d is a leap year!!\n",year)

  }else {
    fmt.Printf("The year %d is not a leap year!!!\n",year)
  }
}
