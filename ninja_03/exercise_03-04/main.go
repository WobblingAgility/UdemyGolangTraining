package main

import "fmt"

func main() {
	i := 1994
	age := 0
	for i <= 2018 {
		fmt.Printf("In %v I was %v\n", i, age)
		i++
		age++
	}
}
