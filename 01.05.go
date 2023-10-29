package main

import "fmt"

type age int

var x age
var y int

func main() {
	fmt.Printf("x = value: %v\ttype: %T\n", x, x)
	x = 42
	fmt.Printf("x = value: %v\n", x)

	y = int(x)
	fmt.Printf("y = value: %v\ttype: %T\n", y, y)
}
