package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)

	xm := make(map[rune]int)
	for _, r := range n {
		xm[r]++
	}
	fmt.Println(xm['2'])
}
