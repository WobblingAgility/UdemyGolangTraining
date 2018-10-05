package main

import "fmt"

func main() {
	x := []int{34, 6, 23, 765, 3, 42, 67, 3, 65, 25}

	for i, v := range x {
		fmt.Printf("%v\t%v\t%T\n", i, v, v)
	}
}
