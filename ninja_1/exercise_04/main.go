package main

import "fmt"

type joe = int

var x joe

func main() {
	fmt.Printf("Variable has the zero value of %v and is type %T\n", x, x)
	x = 42
	fmt.Printf("Variable x = %v\n", x)
}
