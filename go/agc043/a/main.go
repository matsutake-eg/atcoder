package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	iv, _ := strconv.Atoi(scanString())
	return iv
}

func scanString() string {
	sc.Scan()
	return sc.Text()
}

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 100001), 100001*100)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	h, w := scanInt(), scanInt()
	dp := make([][]int, h)
	s := make([]string, h)
	for i := range dp {
		dp[i] = make([]int, w)
		s[i] = scanString()
		for j := range dp[i] {
			if i == 0 && j == 0 {
				if s[i][j] == '#' {
					dp[i][j] = 1
				} else {
					dp[i][j] = 0
				}
			} else {
				dp[i][j] = 101
			}

			if i > 0 {
				x := 0
				if s[i-1][j] == '.' && s[i][j] == '#' {
					x++
				}
				dp[i][j] = min(dp[i][j], dp[i-1][j]+x)
			}
			if j > 0 {
				x := 0
				if s[i][j-1] == '.' && s[i][j] == '#' {
					x++
				}
				dp[i][j] = min(dp[i][j], dp[i][j-1]+x)
			}
		}
	}
	fmt.Println(dp[h-1][w-1])
}
