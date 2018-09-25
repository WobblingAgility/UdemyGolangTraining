package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Printf("The zero values in Golang are:\na. %T = %v\nb. %T = %v\nc. %T = %v\n", x, x, y, y, z, z)
}
