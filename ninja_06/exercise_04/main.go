package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println(p.first)
	fmt.Println(p.last)
}

func main() {
	jd := person{
		"John",
		"Doe",
		23,
	}
	jd.speak()
}
