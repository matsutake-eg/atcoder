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

var (
	par  []int
	rank []int
)

func initTree(x int) {
	par = make([]int, x+1)
	rank = make([]int, x+1)
	for i := range par {
		par[i] = i
	}
}

func find(x int) int {
	if par[x] == x {
		return x
	}
	par[x] = find(par[x])
	return par[x]
}

func union(x, y int) {
	px, py := find(x), find(y)
	if px == py {
		return
	}
	if rank[px] < rank[py] {
		par[px] = py
		return
	} else if rank[px] == rank[py] {
		rank[px]++
	}
	par[py] = px
}

func main() {
	n, m := scanInt(), scanInt()
	initTree(n)
	for i := 0; i < m; i++ {
		a, b := scanInt(), scanInt()
		union(a, b)
	}
	xm := make(map[int]bool)
	for i := 1; i <= n; i++ {
		xm[find(i)] = true
	}
	fmt.Println(len(xm) - 1)
}
