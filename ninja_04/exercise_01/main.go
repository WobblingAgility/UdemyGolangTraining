package main

import "fmt"

func main() {
	a := [5]int{1, 5, 3, 67, 3}

	for i := 0; i < len(a); i++ {
		fmt.Printf("%v is type %T\n", a[i], a[i])
	}
}
