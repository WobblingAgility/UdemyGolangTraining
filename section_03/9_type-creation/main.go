package main

import "fmt"

type hotdog int

var b hotdog

func main() {
	a := 43
	b = 75
	fmt.Printf("Variable \"a\" = %v and is type %T.\n", a, a)
	fmt.Printf("Variable \"b\" = %v and is type %T.\n", b, b)

	a = int(b)
	fmt.Printf("Variable \"a\" = %v and is now equal to \"b\" but type %T.\n", a, a)
}
