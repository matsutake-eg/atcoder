package main

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	n %= k
	for {
		bn := n
		n = min(n, abs(n-k))
		if n == bn {
			fmt.Println(n)
			return
		}
	}
}
