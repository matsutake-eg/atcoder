package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const inf = 1000000000

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	iv, _ := strconv.Atoi(sc.Text())
	return iv
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	n, m := readInt(), readInt()
	dp := make([]int, 1<<uint(n))
	for i := 1; i < len(dp); i++ {
		dp[i] = inf
	}

	for i := 0; i < m; i++ {
		a, b := readInt(), readInt()
		c := 0
		for j := 0; j < b; j++ {
			c += 1 << uint(readInt()-1)
		}
		for j := 0; j < len(dp)-1; j++ {
			dp[j|c] = min(dp[j|c], dp[j]+a)
		}
	}
	if v := dp[len(dp)-1]; v == inf {
		fmt.Println(-1)
	} else {
		fmt.Println(v)
	}
}
