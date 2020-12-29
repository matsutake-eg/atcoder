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
	n, a, b := scanInt(), scanInt(), scanInt()
	x := make([]int, n)
	for i := range x {
		x[i] = scanInt()
	}

	ans := 0
	for i := 0; i < n-1; i++ {
		ans += min((x[i+1]-x[i])*a, b)
	}
	fmt.Println(ans)
}
