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
	seen  []bool
	ans   int
	found bool
)

func dfs(from, to int) {
	seen[to] = true
	for t := range graph[to] {
		if t != from && seen[t] {
			found = true
		}
		if seen[t] {
			continue
		}
		dfs(to, t)
	}
}

func main() {
	n, m := scanInt(), scanInt()
	initGraph(n)
	for i := 0; i < m; i++ {
		v1, v2 := scanInt(), scanInt()
		add(v1, v2)
		add(v2, v1)
	}

	seen = make([]bool, n+1)
	ans = 0
	for i := 1; i <= n; i++ {
		if seen[i] {
			continue
		}
		found = false
		dfs(0, i)
		if !found {
			ans++
		}
	}
	fmt.Println(ans)
}
