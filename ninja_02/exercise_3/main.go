package main

import "fmt"

const (
	t1 = 43
	t2 = "Testing"
	t3 = true
)

const (
	ut1 int32  = 43
	ut2 string = "Testing"
	ut3 bool   = true
)

func main() {
	fmt.Printf("%v is typed %T\n", t1, t1)
	fmt.Printf("%v is typed %T\n", t2, t2)
	fmt.Printf("%v is typed %T\n", t3, t3)
	fmt.Printf("%v is untyped %T\n", ut1, ut1)
	fmt.Printf("%v is untyped %T\n", ut2, ut2)
	fmt.Printf("%v is untyped %T\n", ut3, ut3)
}
