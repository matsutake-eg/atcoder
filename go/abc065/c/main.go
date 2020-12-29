package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	var ans int
	if v := abs(n - m); v == 0 {
		ans = 2
	} else if v == 1 {
		ans = 1
	} else {
		fmt.Println(0)
		return
	}
	const mod = 1000000007
	for i := 2; i <= n; i++ {
		ans *= i
		ans %= mod
	}
	for i := 2; i <= m; i++ {
		ans *= i
		ans %= mod
	}
	fmt.Println(ans)
}
