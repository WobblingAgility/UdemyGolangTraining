package main

import "fmt"

func main() {
	foo()
	fmt.Println()
	bar("Jonny")
	s1 := woo("Jonny")
	fmt.Println(s1)
}

func foo() {
	fmt.Println("Hello from foo!")
}

func bar(s string) {
	fmt.Println("Hello", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s)
}
