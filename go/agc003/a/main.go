package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	xm := make(map[rune]bool)
	for _, r := range s {
		xm[r] = true
	}
	if xm['N'] == xm['S'] && xm['E'] == xm['W'] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
