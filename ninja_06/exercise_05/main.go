package main

import "fmt"

type shape interface {
	area() float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return 2 * s.length
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func info(s shape) float64 {
	return s.area()
}

func main() {
	s := square{
		length: 32,
	}
	c := circle{
		radius: 52,
	}

	fmt.Println(info(s))
	fmt.Println(info(c))
}
