package main

import "fmt"

func main() {
	foo(1, 2, 3, 4, 5, 6, 7, 8, 9)
}

func foo(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
