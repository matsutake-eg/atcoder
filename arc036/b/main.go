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
	n := scanInt()
	h := make([]int, n)
	for i := range h {
		h[i] = scanInt()
	}

	u := make([]int, n)
	for i := 1; i < n; i++ {
		if h[i] > h[i-1] {
			u[i] = u[i-1] + 1
		}
	}
	d := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		if h[i] > h[i+1] {
			d[i] = d[i+1] + 1
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, u[i]+d[i]+1)
	}
	fmt.Println(ans)
}
