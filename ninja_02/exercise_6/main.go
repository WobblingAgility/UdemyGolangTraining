package main

import "fmt"

const (
	y1 = 2014 + iota
	y2 = 2014 + iota
	y3 = 2014 + iota
	y4 = 2014 + iota
)

func main() {
	fmt.Println(y1, y2, y3, y4)
}
