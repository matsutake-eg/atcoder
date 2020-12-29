package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if b > a {
		a, b = b, a
	}
	if a == b {
		fmt.Println(a * 2)
	} else {
		fmt.Println(a*2 - 1)
	}
}
