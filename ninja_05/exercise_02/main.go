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

	fmt.Println()

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Printf("%v %v's favorite flavor is:\n", v.first, v.last)
		for _, s := range v.icecream {
			fmt.Printf("\t%v\n", s)
		}
	}
}
