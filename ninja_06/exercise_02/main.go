package main

import "fmt"

func main() {
	fmt.Println(foo(1, 2, 3, 4, 5, 6, 7))
	n := []int{1, 2, 4, 5, 6, 7, 43}
	fmt.Println(bar(n))
}

func foo(n ...int) int {
	var sum int
	for _, v := range n {
		sum += v
	}
	return sum
}

func bar(n []int) int {
	var sum int
	for _, v := range n {
		sum += v
	}
	return sum
}
