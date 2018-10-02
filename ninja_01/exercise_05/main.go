package main

import "fmt"

type joe = int

var x joe
var y int

func main() {
	fmt.Printf("Variable x has the zero value of %v and is type %T\n", x, x)
	x = 42
	fmt.Printf("Variable x = %v\n", x)

	y = int(x)
	fmt.Printf("Variable y has the value of x \"%v\" but is type %T\n", y, y)
}
