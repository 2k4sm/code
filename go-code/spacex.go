package main

import "fmt"

func main() {
	const (
		speed float64 = 100800   //km/hr
		dist  float64 = 96300000 //km
	)
	var (
		time float64
		hr   float64
		day  float64
	)

	time = dist / speed
	hr = time / 60
	day = hr / 24
	fmt.Printf("The time taken to reach there is: %v minutes.\n", time)
	fmt.Printf("The time taken to reach there is: %v hours.\n", hr)
	fmt.Printf("The time taken to reach there is: %v days.\n", day)

}
