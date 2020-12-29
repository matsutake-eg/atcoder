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
	a, b, m := scanInt(), scanInt(), scanInt()
	amin := math.MaxInt64
	as := make([]int, a)
	for i := range as {
		as[i] = scanInt()
		amin = min(amin, as[i])
	}
	bmin := math.MaxInt64
	bs := make([]int, b)
	for i := range bs {
		bs[i] = scanInt()
		bmin = min(bmin, bs[i])
	}

	ans := amin + bmin
	for i := 0; i < m; i++ {
		x, y, c := scanInt(), scanInt(), scanInt()
		ans = min(ans, as[x-1]+bs[y-1]-c)
	}
	fmt.Println(ans)
}
