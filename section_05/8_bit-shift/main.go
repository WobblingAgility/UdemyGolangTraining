package main

import "fmt"

const (
	_   = iota
	kbi = 1 << (iota * 10)
	mbi = 1 << (iota * 10)
	gbi = 1 << (iota * 10)
)

func main() {
	fmt.Printf("Decimal\t\t\tBinary\n")

	fmt.Println("\nBit shift test")
	x := 4
	fmt.Printf("%d\t\t\t%b\n", x, x)
	y := x << 1
	fmt.Printf("%d\t\t\t%b\n", y, y)
	y = x >> 1
	fmt.Printf("%d\t\t\t%b\n", y, y)

	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Println("\nkb to mb to gb with multiplication")
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t\t%b\n", gb, gb)

	fmt.Println("\nkb to mb to gb with bit shift and iota")
	fmt.Printf("%d\t\t\t%b\n", kbi, kbi)
	fmt.Printf("%d\t\t\t%b\n", mbi, mbi)
	fmt.Printf("%d\t\t\t%b\n", gbi, gbi)
}
