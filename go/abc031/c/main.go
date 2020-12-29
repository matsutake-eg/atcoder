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

func main() {
	n := scanInt()
	a := make([]int, n)
	for i := range a {
		a[i] = scanInt()
	}

	ans := math.MinInt64
	for i := range a {
		amaxs := make([]int, n)
		tmaxs := make([]int, n)
		for i := range amaxs {
			amaxs[i] = math.MinInt64
			tmaxs[i] = math.MinInt64
		}
		amax := math.MinInt64
		for j := range a {
			if i == j {
				continue
			}

			var start, end int
			if i < j {
				start, end = i, j
			} else {
				start, end = j, i
			}
			tsum := 0
			asum := 0
			for k := start; k <= end; k += 2 {
				tsum += a[k]
				if k+1 <= end {
					asum += a[k+1]
				}
			}
			amaxs[j] = asum
			tmaxs[j] = tsum
			amax = max(amax, asum)
		}

		for j := 0; j < n; j++ {
			if amaxs[j] == amax {
				ans = max(ans, tmaxs[j])
				break
			}
		}
	}
	fmt.Println(ans)
}
