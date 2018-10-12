package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	for i, _ := range m {
		fmt.Println("These are the values for", i)
		for j, v := range m[i] {
			fmt.Printf("%v\t%v\n", j, v)
		}
	}

	m["Todd"] = []string{`Golang`, `Teaching`, `Quotes`}
	fmt.Printf("\nI've added to the list!\n\n")

	for i, _ := range m {
		fmt.Println("These are the values for", i)
		for j, v := range m[i] {
			fmt.Printf("%v\t%v\n", j, v)
		}
	}

	delete(m, `bond_james`)
	fmt.Printf("\nWe lost someone!\n\n")

	for i, _ := range m {
		fmt.Println("These are the values for", i)
		for j, v := range m[i] {
			fmt.Printf("%v\t%v\n", j, v)
		}
	}
}
