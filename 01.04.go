package main

import "fmt"

type age int

var x age

func main() {
	fmt.Printf("x = value: %v\ttype: %T\n", x, x)

	x = 42

	fmt.Printf("x = value: %v\n", x)
}
