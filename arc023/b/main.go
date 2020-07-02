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
	r, c, d := scanInt(), scanInt(), scanInt()
	a := make([][]int, r)
	for i := range a {
		a[i] = make([]int, c)
		for j := range a[i] {
			a[i][j] = scanInt()
		}
	}

	ans := math.MinInt64
	for i := 0; i < min(d+1, r); i++ {
		for j := (i + d) % 2; j < min(d+1-i, c); j += 2 {
			ans = max(ans, a[i][j])
		}
	}
	fmt.Println(ans)
}
