package main

import "fmt"

func main() {
	year := 1999
	currentYear := 2023

	for {
		fmt.Println(year)

		if year == currentYear {
			break
		}

		year++
	}
}
