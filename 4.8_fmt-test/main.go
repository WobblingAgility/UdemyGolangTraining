package main

import "fmt"

var y int = 42

func main() {
	fmt.Println(y)
	fmt.Printf("Variable y is type %T.\n", y)
	fmt.Printf("Variable y is %b in binary.\n", y)
	fmt.Printf("Variable y is %x in hexidecimal.\n", y)
	fmt.Printf("Variable y is %#x in formatted hexidecimal.\n", y)
}
