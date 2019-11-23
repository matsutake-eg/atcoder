package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const mod = 1000000007

var fac, finv, inv []int
var wr = bufio.NewWriter(os.Stdout)

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

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	initCmb(max(n-k+1, k-1))
	for i := 1; i <= k; i++ {
		wr.WriteString(strconv.Itoa(cmb(n-k+1, i)*cmb(k-1, i-1)%mod) + "\n")
	}
	wr.Flush()
}
