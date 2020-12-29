package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	m := make(map[rune]bool)
	for _, r := range s {
		m[r] = true
	}
	if len(m) == len(s) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
