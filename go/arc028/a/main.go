package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	for {
		n -= a
		if n <= 0 {
			fmt.Println("Ant")
			return
		}
		n -= b
		if n <= 0 {
			fmt.Println("Bug")
			return
		}
	}
}
