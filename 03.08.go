package main

import "fmt"

func main() {
	age := 900

	switch {
	case age >= 500:
		fmt.Println("You are so old!")
	case age > 0:
		fmt.Println("You are so young!")
	default:
		fmt.Println("Are you sure?")
	}
}
