package main

import (
        "fmt"
        "math/rand"
)

func main() {

  fmt.Println("Spaceline            Days     Trip-Type     Price  ")
  fmt.Println("=================================================")
  list := []string {"SpaceX","SpaceAdventures","Virgin Galactic"}
  fmt.Println(list)

  speed := rand.Intn(50)+1 //km/s
  dist  := 62100000  //km
  time_hr:= dist/speed/3600/24 //sec divided for hour then for day
  fmt.Println(time_hr)

  

}
