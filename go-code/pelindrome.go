package main

import ("fmt"
        "math"
)

func powInt(x, y int) int {
    return int(math.Pow(float64(x), float64(y)))
}

func main(){
  var (
    num1 int
    num2 int
    temp int = 0
  )

  fmt.Println("Enter the digits present in the number:")
  fmt.Scanln(&num2)
  fmt.Println("Enter the number:")
  fmt.Scanln(&num1)

  var dup int =num2-1
  //fmt.Println(dup)
  i:=powInt(10,dup)

  for i!=0{
    temp+=(num1%10)*i
    num1=num1/10
    i/=10
  }
  fmt.Printf("The pelindrome of the given %d digit number is %d ",num2,temp)
}
