package main

import "fmt"

type person struct {
	first    string
	last     string
	icecream []string
}

func main() {
	p1 := person{
		first:    `mr`,
		last:     `hanesbero`,
		icecream: []string{`chocolate`, `vanila`},
	}
	p2 := person{
		first:    `Johny`,
		last:     `Cash`,
		icecream: []string{`cash`, `cherry`},
	}

	fmt.Printf("%v %v's favorite icecream flavors are:\n", p1.first, p1.last)
	for _, v := range p1.icecream {
		fmt.Printf("\t%v\n", v)
	}
	fmt.Printf("%v %v's favorite icecream flavors are:\n", p2.first, p2.last)
	for _, v := range p2.icecream {
		fmt.Printf("\t%v\n", v)
	}
}
