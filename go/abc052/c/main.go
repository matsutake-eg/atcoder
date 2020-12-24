package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	xm := make(map[int]int)
	for i := 2; i <= n; i++ {
		x := i
		for j := 2; j <= i; j++ {
			for x%j == 0 {
				xm[j]++
				x /= j
			}
		}
	}

	const mod = 1000000007
	ans := 1
	for _, v := range xm {
		ans *= v + 1
		ans %= mod
	}
	fmt.Println(ans)
}
