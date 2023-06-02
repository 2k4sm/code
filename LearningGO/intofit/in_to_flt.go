package main

import(
  f "fmt"
  r "reflect"

)
func main(){

  var id = 13
  f.Println("id is:",id)
  f.Println("id type:",r.TypeOf(id))


  n_id:= float64(id)
  f.Println("id is:",n_id)
  f.Println("id type is:",r.TypeOf(n_id))

}
