package main

import "fmt"

func main() {
	equal := 2 == 2
	diff := 2 != 2
	lesser := 2 < 2
	greater := 2 > 2
	lesser_equal := 2 <= 2
	greater_equal := 2 >= 2

	fmt.Printf("2 == 2: %v\n", equal)
	fmt.Printf("2 != 2: %v\n", diff)
	fmt.Printf("2 < 2: %v\n", lesser)
	fmt.Printf("2 > 2: %v\n", greater)
	fmt.Printf("2 <= 2: %v\n", lesser_equal)
	fmt.Printf("2 >= 2: %v\n", greater_equal)
}
