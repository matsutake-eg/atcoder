package main

import "fmt"

func main() {
	var (
		a, b int
		op   string
	)
	fmt.Scan(&a, &op, &b)

	if op == "+" {
		fmt.Println(a + b)
	} else {
		fmt.Println(a - b)
	}
}
