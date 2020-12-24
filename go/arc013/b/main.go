package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	c := scanInt()
	xs := make([][]int, c)
	mn, mm, ml := 0, 0, 0
	for i := range xs {
		xs[i] = make([]int, 3)
		for j := range xs[i] {
			xs[i][j] = scanInt()
		}
		sort.Ints(xs[i])
		mn = max(mn, xs[i][0])
		mm = max(mm, xs[i][1])
		ml = max(ml, xs[i][2])
	}
	fmt.Println(mn * mm * ml)
}
