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
	n := scanInt()
	mx := math.MinInt64
	mn := math.MaxInt64
	for i := 0; i < n; i++ {
		a := scanInt()
		mx = max(mx, a)
		mn = min(mn, a)
	}
	fmt.Println(mx - mn)
}
