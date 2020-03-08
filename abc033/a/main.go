package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)

	xm := make(map[rune]bool)
	for _, r := range n {
		xm[r] = true
	}
	if len(xm) == 1 {
		fmt.Println("SAME")
	} else {
		fmt.Println("DIFFERENT")
	}
}
