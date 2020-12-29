package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	iv, _ := strconv.Atoi(sc.Text())
	return iv
}

func init() {
	sc.Split(bufio.ScanWords)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	n, m, k := scanInt(), scanInt(), scanInt()
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = scanInt() + a[i-1]
	}
	b := make([]int, m+1)
	for i := 1; i <= m; i++ {
		b[i] = scanInt() + b[i-1]
	}

	ans := 0
	p := len(b) - 1
	for i := range a {
		if a[i] > k {
			break
		}
		for j := p; j >= 0; j-- {
			if a[i]+b[j] <= k {
				p = j
				ans = max(ans, i+j)
				break
			}
		}
	}
	fmt.Println(ans)
}
