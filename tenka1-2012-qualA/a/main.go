package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	xs := make([]int, 46)
	xs[0] = 1
	xs[1] = 1
	for i := 2; i < len(xs); i++ {
		xs[i] = xs[i-1] + xs[i-2]
	}
	fmt.Println(xs[n])
}
