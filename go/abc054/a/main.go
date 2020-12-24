package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if a == 1 {
		a += 13
	}
	if b == 1 {
		b += 13
	}

	if a > b {
		fmt.Println("Alice")
	} else if a < b {
		fmt.Println("Bob")
	} else {
		fmt.Println("Draw")
	}
}
