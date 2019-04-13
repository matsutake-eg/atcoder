package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	sum := 0
	for i := 0; i < 2; i++ {
		if a > b {
			sum += a
			a--
		} else {
			sum += b
			b--
		}
	}

	fmt.Println(sum)
}
