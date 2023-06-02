package main

import "fmt"

func main() {
  var publisher,writer,artist,title string
  publisher = "DizzyBooks Publishing Inc."
  writer = "Tracey Hatchet"
  artist = "Jewel Tampson"
  title = "Mr. GoToSleep"
  var year,pageNumber int
  year = 1997
  pageNumber = 14
  var grade float32
  grade = 6.5
  fmt.Println(title, "written by", writer, "drawn by", artist)
  fmt.Println("Published by",publisher,"Year published",year,"Total pages",pageNumber)
  fmt.Println("Grade of the boook is",grade)



}
