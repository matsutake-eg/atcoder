package main

import "fmt"

const mod = 1000000007

var (
	n, k int
	c    [][]int
)

func initCmb() {
	max := n - k + 1
	if v := k - 1; v > max {
		max = v
	}

	c = make([][]int, max+1)
	for i := range c {
		c[i] = make([]int, i+1)
		c[i][0], c[i][len(c[i])-1] = 1, 1
		for j := 1; j < len(c[i])-1; j++ {
			c[i][j] = (c[i-1][j-1] + c[i-1][j]) % mod
		}
	}
}

func cmb(a, b int) int {
	if b > a {
		return 0
	}
	return c[a][b]
}

func main() {
	fmt.Scan(&n, &k)

	initCmb()
	for i := 1; i <= k; i++ {
		fmt.Println(cmb(n-k+1, i) * cmb(k-1, i-1) % mod)
	}
}
