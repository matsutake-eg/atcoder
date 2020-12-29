package main

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	var n int
	fmt.Scan(&n)
	b := make([]int, n-1)
	for i := range b {
		fmt.Scan(&b[i])
	}

	ans := b[0] + b[len(b)-1]
	for i := range b[:len(b)-1] {
		ans += min(b[i], b[i+1])
	}
	fmt.Println(ans)
}
