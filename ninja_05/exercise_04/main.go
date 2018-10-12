package main

import "fmt"

func main() {
	s := struct {
		first    string
		last     string
		treasure map[string]int
		crew     []string
	}{
		first: `Captain`,
		last:  `Jack`,
		treasure: map[string]int{
			`first`:  304,
			`secone`: 394,
			`third`:  342,
		},
		crew: []string{`John`, `Mike`, `Bob`},
	}
	fmt.Println(s)
}
