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

func main() {
	n, a, b := scanInt(), scanInt(), scanInt()
	mx := math.MinInt64
	mn := math.MaxInt64
	sum := 0
	for i := 0; i < n; i++ {
		s := scanInt()
		mx = max(mx, s)
		mn = min(mn, s)
		sum += s
	}

	if mx-mn == 0 {
		fmt.Println(-1)
		return
	}

	p := float64(b) / float64(mx-mn)
	q := (float64(a*n) - p*float64(sum)) / float64(n)
	fmt.Println(p, q)
}
