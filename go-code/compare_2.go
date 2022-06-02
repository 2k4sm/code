package main

import "fmt"

func main(){

  var a int64
  var b int64

  fmt.Println("Enter your first integer:")
  fmt.Scanln(&a)

  fmt.Println("Enter your second integer:")
  fmt.Scanln(&b)

  if a>b {
    fmt.Println(a,"is the greatest integer.")

  }else {

    fmt.Println(b,"is the greatest integer.")
  }



}
