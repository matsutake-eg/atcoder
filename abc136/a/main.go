package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	r := c - (a - b)
	if r < 0 {
		fmt.Println(0)
		return
	}
	fmt.Println(r)
}
