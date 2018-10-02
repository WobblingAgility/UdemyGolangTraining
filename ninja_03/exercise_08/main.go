package main

import "fmt"

func main() {
	x := "go"
	switch {
	case x == "no go":
		fmt.Println("no go")
	case x == "go":
		fmt.Println("go")
	default:
		fmt.Println("go?")
	}
}
