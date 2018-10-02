package main

import "fmt"

func main() {
	if x := 42; x != 2 {
		fmt.Println(x)
	}

	if x := "test"; x == "Test" {
		fmt.Println("Test")
	} else if x == "test" {
		fmt.Println("test")
	}
}
