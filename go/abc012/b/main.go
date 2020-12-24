package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	h := n / 3600
	n %= 3600
	m := n / 60
	n %= 60
	fmt.Printf("%02v:%02v:%02v\n", h, m, n)
}
