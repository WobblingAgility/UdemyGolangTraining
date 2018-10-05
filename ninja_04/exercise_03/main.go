package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range x {
		fmt.Printf("%v\t%v\t%T\n", i, v, v)
	}

	fmt.Println(x[0:5])
	fmt.Println(x[5:10])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
}
