package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	xs := make([]int, 0, n-1)
	for i := 1; i <= n; i++ {
		if i != m {
			xs = append(xs, i)
		}
	}
	fmt.Println(xs[0])
}
