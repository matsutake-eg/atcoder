package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const dft = -1

var match = [9]int{2, 5, 5, 4, 5, 6, 3, 7, 6}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([]int, m)
	for i := range a {
		fmt.Scan(&a[i])
	}

	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = dft
	}
	dp[0] = 0

	for i := range dp {
		if dp[i] == dft {
			continue
		}
		for _, v := range a {
			mv := match[v-1]
			if i+mv > n {
				continue
			}
			if dp[i+mv] == dft || dp[i]+1 > dp[i+mv] {
				dp[i+mv] = dp[i] + 1
			}
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	wr := bufio.NewWriter(os.Stdout)
	for i := n; i > 0; {
		for _, v := range a {
			mv := match[v-1]
			if i-mv >= 0 && dp[i-mv] == dp[i]-1 {
				wr.WriteString(strconv.Itoa(v))
				i -= mv
				break
			}
		}
	}
	wr.WriteString("\n")
	wr.Flush()
}
