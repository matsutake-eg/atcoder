package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	const mod = 1000000007
	ans := 1
	for i := 1; i <= n; i++ {
		ans *= i
		ans %= mod
	}
	fmt.Println(ans)
}
