package main

import "fmt"

func main() {
	foo("This message was defered", true)
	foo("This message was not", false)
}

func foo(s string, d bool) {
	if d {
		defer fmt.Println(s)
	} else {
		fmt.Println(s)
	}
	fmt.Println("The function ended here")
}
