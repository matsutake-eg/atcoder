package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	xm := make(map[rune]int)
	for _, r := range s {
		xm[r]++
	}
	fmt.Println(xm['1'])
}
