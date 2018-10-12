package main

import "fmt"

func main() {
	foo(1, 2, 3, 4, 5, 6, 7, 8, 9)
}

// This is a variadic parameter. By definition it is a functions whos parameter takes zero or more values of a specific type
// ... before a type "unfurls" the type so that the fuction takes unlimited of that type
func foo(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position,", i, "we are now adding,", v, "to the total thich is now", sum)
	}
	fmt.Println("the total is,", sum)
}
