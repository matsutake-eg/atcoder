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
	n := scanInt()
	a := make([]int, n)
	for i := range a {
		a[i] = scanInt()
	}
	sort.Ints(a)

	m := (n + 1) / 2
	x := a[m]
	if n%2 == 0 {
		x -= a[m-1]
	} else {
		x += a[m-1]
	}
	for i := 0; i < m-1; i++ {
		x -= a[i] * 2
	}
	for i := m + 1; i < n; i++ {
		x += a[i] * 2
	}
	if n%2 == 0 {
		fmt.Println(x)
		return
	}

	m = (n - 1) / 2
	y := -a[m] - a[m-1]
	for i := 0; i < m-1; i++ {
		y -= a[i] * 2
	}
	for i := m + 1; i < n; i++ {
		y += a[i] * 2
	}
	fmt.Println(max(x, y))
}
