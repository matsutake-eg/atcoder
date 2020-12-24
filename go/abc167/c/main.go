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

func main() {
	n, m, x := scanInt(), scanInt(), scanInt()
	c := make([]int, n)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		c[i] = scanInt()
		a[i] = make([]int, m)
		for j := range a[i] {
			a[i][j] = scanInt()
		}
	}

	ans := math.MaxInt64
	for i := 1; i < 1<<uint(n); i++ {
		csum := 0
		asum := make([]int, m)
		for j := 0; j < n; j++ {
			if b := 1 << uint(j); i&b == b {
				csum += c[j]
				for k := range a[j] {
					asum[k] += a[j][k]
				}
			}
		}

		isOver := true
		for _, av := range asum {
			if av < x {
				isOver = false
				break
			}
		}
		if isOver {
			ans = min(ans, csum)
		}
	}
	if ans == math.MaxInt64 {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}
