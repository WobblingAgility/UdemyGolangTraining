package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mp := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	ss := [][]string{jb, mp}
	fmt.Println(ss)

	for i, _ := range ss {
		for j, _ := range ss[i] {
			fmt.Println(ss[i][j])
		}
	}
}
