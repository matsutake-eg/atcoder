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

var graph map[int]map[int]bool

func initGraph(x int) {
	graph = make(map[int]map[int]bool, x)
}

func add(from, to int) {
	if _, ok := graph[from]; !ok {
		graph[from] = make(map[int]bool)
	}
	graph[from][to] = true
}

var (
	cnt []int
	ans []int
)

func dfs(v, pv int) {
	ans[v] = ans[pv] + cnt[v]
	for t := range graph[v] {
		if ans[t] >= 0 {
			continue
		}
		dfs(t, v)
	}
}

func main() {
	n, q := scanInt(), scanInt()
	initGraph(n)
	for i := 0; i < n-1; i++ {
		a, b := scanInt(), scanInt()
		add(a, b)
		add(b, a)
	}

	cnt = make([]int, n+1)
	for i := 0; i < q; i++ {
		p, x := scanInt(), scanInt()
		cnt[p] += x
	}

	ans = make([]int, n+1)
	for i := 1; i <= n; i++ {
		ans[i] = -1
	}

	dfs(1, 0)
	var wr = bufio.NewWriter(os.Stdout)
	for i := 1; i <= n; i++ {
		wr.WriteString(strconv.Itoa(ans[i]))
		if i < n {
			wr.WriteString(" ")
		} else {
			wr.WriteString("\n")
		}
	}
	wr.Flush()
}
