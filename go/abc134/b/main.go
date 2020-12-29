package main

import "fmt"

func main() {
	var n, d int
	fmt.Scan(&n, &d)

	c := 1 + d*2
	fmt.Println((n + c - 1) / c)
}
