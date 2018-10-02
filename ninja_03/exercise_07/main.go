package main

import "fmt"

func main() {
	x := "go"
	if x == "no go" {
		fmt.Println("no go")
	} else if x == "go" {
		fmt.Println("go")
	} else {
		fmt.Println("go?")
	}
}
