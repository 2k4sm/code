package main

import "fmt"

func main() {
	fmt.Println("This is a simple switch-case statement...")

	var light string
	fmt.Printf("Enter the state of the light in the room:")
	fmt.Scanf("%s", &light)

	switch light {
	case "off":
		fmt.Println("The light is off you can simply sleep...")
	case "on":
		fmt.Println("The light is on you need to turn it off then sleep...")
	default:
		fmt.Println("just sleep! goodnight...")
	}
}
