package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])

	v, ok := m["test"]
	fmt.Println(v)
	fmt.Println(ok)

	//Add to a map
	m["todd"] = 33

	fmt.Println()
	for k, v := range m {
		fmt.Println(k, v)
	}
}
