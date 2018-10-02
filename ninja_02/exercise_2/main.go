package main

import "fmt"

func main() {
	a := 4 == 3
	b := 4 <= 3
	c := 4 >= 3
	d := 4 != 3
	e := 4 < 3
	f := 4 > 3

	fmt.Printf("Is 4 equal to 3? %v\n", a)
	fmt.Printf("Is 4 less than or equal to 3? %v\n", b)
	fmt.Printf("Is 4 greater than or equal to 3? %v\n", c)
	fmt.Printf("Is 4 not equal to 3? %v\n", d)
	fmt.Printf("Is 4 less to 3? %v\n", e)
	fmt.Printf("Is 4 greater to 3? %v\n", f)
}
