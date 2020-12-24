package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if b < a {
		fmt.Println(a - 1)
		return
	}
	fmt.Println(a)
}
