package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	xs := make([]int, n+1)
	for i := range xs {
		xs[i] = 1
	}
	for i := 2; i <= n; i++ {
		for j := 1; i*j <= n; j++ {
			xs[i*j]++
		}
	}
	ans := 0
	for i := 1; i <= n; i++ {
		ans += i * xs[i]
	}
	fmt.Println(ans)
}
