package main

import (
	"fmt"
)

const mod = 1000000007

func pow(x, y int) int {
	if y == 0 {
		return 1
	} else if y%2 == 0 {
		x2 := pow(x, y/2)
		return x2 * x2 % mod
	} else {
		return x * pow(x, y-1) % mod
	}
}

func cmb(x, y int) int {
	if y > x {
		return 0
	}

	nu, de := 1, 1
	for i := 0; i < y; i++ {
		nu = nu * (x - i) % mod
		de = de * (i + 1) % mod
	}
	return nu * pow(de, mod-2) % mod
}

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	ans := pow(2, n) - 1
	ans -= cmb(n, a)
	ans -= cmb(n, b)
	ans = (ans%mod + mod) % mod
	fmt.Println(ans)
}
