package main

import (
        "fmt"
        "math/rand"
)

func main() {
  var trip_type string
  fmt.Printf("Enter the type of trip you are making(One-way or Round-trip):")
  fmt.Scanf("%s",&trip_type)

  fmt.Printf("%-30v %-15v %-15v  %v \n","Spaceline","Days","Trip-type","Price")
  fmt.Println("======================================================================")
  var list = []string {"Virgin Galactic","SpaceX","SpaceAdventures"}
  j:=10
  for j!=0 {
    for _,i := range list {
      speed := rand.Intn(15)+16 //km/s
      dist  := 62100000 //km
      time_d := dist/speed/3600/24 //in days
      //var trip_type string
      //fmt.Printf("Enter the type of trip you are making(One-way or Round-trip):")
      //fmt.Scanf("%s",&trip_type)
      var price int
      switch trip_type {
        case "One-way":
          if time_d < 23 {
            price = time_d*4 //million$

          }else {
            price = time_d*10 //million$
          }
        case "Round-trip":
          if time_d < 23 {
            price = time_d*2*4 //million$
            time_d *=2

          }else {
            price = time_d*2*10 //million$
            time_d*=2

          }

      }
      fmt.Printf("%-30v %-15v %-15v  $%v \n",i,time_d,trip_type,price)
    }
    j-=1
  }

}
