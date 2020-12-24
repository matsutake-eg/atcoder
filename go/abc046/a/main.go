package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	xm := make(map[int]bool)
	xm[a] = true
	xm[b] = true
	xm[c] = true
	fmt.Println(len(xm))
}
