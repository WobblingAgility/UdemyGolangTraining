package main

import "fmt"

func main() {
	favSport := "b-ball"

	switch favSport {
	case "t-ball":
		fmt.Println("Terible favorite sport")
	case "a-ball":
		fmt.Println("Terible favorite sport")
	case "b-ball":
		fmt.Println("Amazing favorite sport")
	default:
		fmt.Println("You need to pick a favorite")
	}
}
