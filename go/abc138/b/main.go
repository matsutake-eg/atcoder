package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	s := 0.0
	var a float64
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		s += 1 / a
	}

	fmt.Println(1 / s)
}
