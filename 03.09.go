package main

import "fmt"

func main() {
	favoriteGame := "zelda"

	switch favoriteGame {
	case "zelda":
		fmt.Println("The legend of zelda is awsome!")
	case "wow":
		fmt.Println("Warlock affliction is the best class!")
	default:
		fmt.Println("IDK this game T-T")
	}
}
