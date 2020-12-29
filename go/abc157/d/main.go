package main

import (
	"bufio"
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
	n, m, k := scanInt(), scanInt(), scanInt()
	initTree(n)
	xm := make(map[int]map[int]bool, n)
	for i := 0; i < m; i++ {
		a, b := scanInt(), scanInt()
		union(a, b)
		if _, ok := xm[a]; !ok {
			xm[a] = make(map[int]bool)
		}
		xm[a][b] = true
		if _, ok := xm[b]; !ok {
			xm[b] = make(map[int]bool)
		}
		xm[b][a] = true
	}
	for i := 0; i < k; i++ {
		c, d := scanInt(), scanInt()
		if _, ok := xm[c]; !ok {
			xm[c] = make(map[int]bool)
		}
		xm[c][d] = true
		if _, ok := xm[d]; !ok {
			xm[d] = make(map[int]bool)
		}
		xm[d][c] = true
	}

	cm := make(map[int]int)
	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fi := find(i)
		cm[fi]++
		for j := range xm[i] {
			if fi == find(j) {
				ans[i]--
			}
		}
	}

	var wr = bufio.NewWriter(os.Stdout)
	for i := 1; i <= n; i++ {
		wr.WriteString(strconv.Itoa(cm[find(i)] + ans[i] - 1))
		if i != n {
			wr.WriteString(" ")
		} else {
			wr.WriteString("\n")
		}
	}
	wr.Flush()
}
