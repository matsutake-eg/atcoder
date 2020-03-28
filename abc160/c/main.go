package main

import (
	"bufio"
	"fmt"
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
	k, n := scanInt(), scanInt()
	a := make([]int, n)
	for i := range a {
		a[i] = scanInt()
	}

	ans := a[n-1] - a[0]
	for i := 1; i < n; i++ {
		ans = min(ans, a[i-1]-(a[i]-k))
	}
	fmt.Println(ans)
}
