package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	iv, _ := strconv.Atoi(sc.Text())
	return iv
}

func init() {
	sc.Split(bufio.ScanWords)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	h, n := readInt(), readInt()
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i] = readInt(), readInt()
	}

	const max = 1000000000
	dp := make([]int, h+1)
	for i := range dp[:h] {
		dp[i] = max
	}

	for i := h; i >= 1; i-- {
		for j := 0; j < n; j++ {
			nh := 0
			if v := i - a[j]; v > 0 {
				nh = v
			}
			dp[nh] = min(dp[nh], dp[i]+b[j])
		}
	}
	fmt.Println(dp[0])
}
