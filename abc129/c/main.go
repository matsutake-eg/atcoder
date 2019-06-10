package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	xs := make([]int, n+1)
	for i := range xs {
		xs[i] = 1
	}
	var (
		sc = bufio.NewScanner(os.Stdin)
		a  int
	)
	for i := 0; i < m; i++ {
		sc.Scan()
		a, _ = strconv.Atoi(sc.Text())
		xs[a] = 0
	}
	var (
		f2     = 1
		f1, f0 = xs[1], xs[0]
	)
	for i := 2; i <= n; i++ {
		f2 = (f1*xs[i-1] + f0*xs[i-2]) % 1000000007
		f1, f0 = f2, f1
	}
	fmt.Println(f2)
}
