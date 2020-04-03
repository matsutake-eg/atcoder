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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	n := scanInt()
	a := make([]int, n)
	for i := range a {
		a[i] = scanInt()
	}

	ans := make([]int, n)
	for i := 1; i < n; i++ {
		ans[i] = math.MaxInt64
	}
	for i := 0; i < n; i++ {
		if ni := i + 1; ni < n {
			ans[ni] = min(ans[ni], ans[i]+abs(a[ni]-a[i]))
		}
		if ni := i + 2; ni < n {
			ans[ni] = min(ans[ni], ans[i]+abs(a[ni]-a[i]))
		}
	}
	fmt.Println(ans[n-1])
}
