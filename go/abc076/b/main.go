package main

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	ans := 1
	for i := n; i > 0; i-- {
		ans += min(ans, k)

	}
	fmt.Println(ans)
}
