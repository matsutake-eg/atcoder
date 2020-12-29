package main

import "fmt"

func main() {
	var a, b string
	fmt.Scan(&a, &b)

	if len(a) > len(b) {
		fmt.Println("GREATER")
	} else if len(a) < len(b) {
		fmt.Println("LESS")
	} else if a > b {
		fmt.Println("GREATER")
	} else if a < b {
		fmt.Println("LESS")
	} else {
		fmt.Println("EQUAL")
	}
}
