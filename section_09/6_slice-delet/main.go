package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x[1])
	fmt.Println(x)
	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}

	y := []int{243, 452, 765, 937}
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
