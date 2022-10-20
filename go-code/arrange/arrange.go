package main

import "fmt"

func main(){
  // OBTAINING THE ARRAY....
  fmt.Println("Enter the size of your array:")
  var size int
  fmt.Scanln(&size)
  var arr = make([]int,size)
  for i:=0;i<size;i++{
    fmt.Printf("Enter the %dth element:",i+1)
    fmt.Scanf("%d",&arr[i])
  }
  //fmt.Println("The resulted array is:",arr)

  //ARRANGING THE ARRAY.....
  temp:=0
  for k:=0;k!=size;k++{
    for j:=0;j<size-1;j++{
      if arr[j]>arr[j+1]{
        temp = arr[j]
        arr[j] = arr[j+1]
        arr[j+1] = temp

      }
      //fmt.Println(arr)
    }
  }
  fmt.Println("The arranged array is:",arr)
}
