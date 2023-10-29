package main

import "fmt"

func main() {
	for x := 10; x <= 100; x++ {
		fmt.Printf("%d %% 4 = %d\n", x, x%4)
	}
}
