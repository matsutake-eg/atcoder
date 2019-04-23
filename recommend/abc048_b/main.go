package main

import "fmt"

func div(a, x int) int {
	if a == 0 {
		return -1
	}
	return (a - 1) / x
}

func main() {
	var a, b, x int
	fmt.Scan(&a, &b, &x)

	fmt.Println(b/x - div(a, x))
}
