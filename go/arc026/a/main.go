package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	if n <= 5 {
		fmt.Println(b * n)
	} else {
		fmt.Println(b*5 + a*(n-5))
	}
}
