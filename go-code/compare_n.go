package main

import "fmt"

func main(){
  fmt.Println("Enter the size of the array:")
  var size int
  fmt.Scanln(&size)
  var arr = make([]int,size)
  for i:=0;i<size;i++ {
    fmt.Printf("Enter the %dth element:",i+1)
    fmt.Scanf("%d",&arr[i])
  }
  fmt.Println("Your array is:",arr)
  temp:=0
  for _, element:= range arr{
    if element>temp{
      temp = element
    }
  }
  fmt.Printf("The Greatest number in the array is: %d",temp)
}
