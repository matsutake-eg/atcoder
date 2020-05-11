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
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		x := scanInt()
		a[i] = a[i-1] + x
	}

	ans := math.MaxInt64
	for i := 1; i < n; i++ {
		ans = min(ans, abs(a[len(a)-1]-a[i]*2))
	}
	fmt.Println(ans)
}
