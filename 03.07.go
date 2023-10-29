package main

import "fmt"

func main() {
	age := 900

	if age >= 500 {
		fmt.Println("You are so old!")
	} else if age > 0 {
		fmt.Println("You are so young!")
	} else {
		fmt.Println("Are you sure?")
	}
}
