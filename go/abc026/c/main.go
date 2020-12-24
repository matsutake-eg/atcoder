package main

import (
	"bufio"
	"fmt"
	"math"
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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

var xm map[int]map[int]bool

func dfs(x int) int {
	switch len(xm[x]) {
	case 0:
		return 1
	case 1:
		for i := range xm[x] {
			return dfs(i)*2 + 1
		}
	}

	mx := math.MinInt64
	mn := math.MaxInt64
	for i := range xm[x] {
		v := dfs(i)
		mx = max(mx, v)
		mn = min(mn, v)
	}
	return mx + mn + 1
}

func main() {
	n := scanInt()
	xm = make(map[int]map[int]bool, n)
	for i := 0; i < n-1; i++ {
		b := scanInt()
		if _, ok := xm[b]; !ok {
			xm[b] = make(map[int]bool)
		}
		xm[b][i+2] = true
	}

	fmt.Println(dfs(1))
}
