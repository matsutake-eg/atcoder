package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	if k == 0 {
		fmt.Println(n * n)
		return
	}

	ans := 0
	for b := 1; b <= n; b++ {
		p, r := n/b, n%b
		ans += p * max(0, b-k)
		ans += max(0, r-k+1)
	}
	fmt.Println(ans)
}
