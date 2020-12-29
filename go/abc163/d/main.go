package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	const mod = 1000000007
	ans := 0
	for i := k; i <= n+1; i++ {
		mn := (i - 1) * i / 2
		mx := (n*2 - i + 1) * i / 2
		ans += mx - mn + 1
		ans %= mod
	}
	fmt.Println(ans)
}
