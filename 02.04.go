package main

import "fmt"

func main() {
	age := 900

	fmt.Printf("dec: %d\t bin: %b\t hex: %#x\n", age, age, age)

	newAge := age << 1

	fmt.Printf("dec: %d\t bin: %b\t hex: %#x\n", newAge, newAge, newAge)
}
