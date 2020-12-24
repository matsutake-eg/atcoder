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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	n := scanInt()
	t := make([]int, n)
	for i := range t {
		t[i] = scanInt()
	}

	ans := math.MaxInt64
	for i := 1; i < 1<<uint(n); i++ {
		sum1, sum2 := 0, 0
		for j := 0; j < n; j++ {
			if b := 1 << uint(j); i&b == b {
				sum1 += t[j]
			} else {
				sum2 += t[j]
			}
		}
		ans = min(ans, max(sum1, sum2))
	}
	fmt.Println(ans)
}
