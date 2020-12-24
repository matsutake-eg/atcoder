package main

import "fmt"

func main() {
	var n, m, c int
	fmt.Scan(&n, &m, &c)

	bs := make([]int, m)
	for i := range bs {
		fmt.Scan(&bs[i])
	}

	count := 0
	for i := 0; i < n; i++ {
		var x int
		sum := c
		for j := 0; j < m; j++ {
			fmt.Scan(&x)
			sum += x * bs[j]
		}
		if sum > 0 {
			count++
		}
	}
	fmt.Println(count)
}
