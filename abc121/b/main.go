package main

import "fmt"

func main() {
	var n, m, c int
	fmt.Scan(&n, &m, &c)
	bs := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&bs[i])
	}

	count := 0
	var x int
	for i := 0; i < n; i++ {
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
