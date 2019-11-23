package main

import "fmt"

const mod = 1000000007

var fac, finv, inv []int

func initCmb(x int) {
	fac, finv, inv = make([]int, x+1), make([]int, x+1), make([]int, x+1)
	fac[0], fac[1] = 1, 1
	finv[0], finv[1] = 1, 1
	inv[1] = 1
	for i := 2; i <= x; i++ {
		fac[i] = fac[i-1] * i % mod
		inv[i] = mod - inv[mod%i]*(mod/i)%mod
		finv[i] = finv[i-1] * inv[i] % mod
	}
}

func cmb(x, y int) int {
	if y > x {
		return 0
	}
	return fac[x] * (finv[y] * finv[x-y] % mod) % mod
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	if (x+y)%3 != 0 {
		fmt.Println(0)
		return
	}

	a := (2*y - x) / 3
	b := (2*x - y) / 3
	if a < 0 || b < 0 {
		fmt.Println(0)
		return
	}

	initCmb(a + b)
	fmt.Println(cmb(a+b, a))
}
