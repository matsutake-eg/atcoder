package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	iv, _ := strconv.Atoi(sc.Text())
	return iv
}

var n int
var G map[int]map[int]bool

func initGraph() {
	G = make(map[int]map[int]bool, n)
}

func add(from, to int) {
	if _, ok := G[from]; !ok {
		G[from] = make(map[int]bool)
	}
	G[from][to] = true
}

var seen map[int]bool
var ba, bb int

func dfs(v int) {
	seen[v] = true
	for t := range G[v] {
		if seen[t] {
			continue
		}
		if (v == ba && t == bb) || (v == bb && t == ba) {
			continue
		}
		dfs(t)
	}
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	n = readInt()
	m := readInt()
	initGraph()

	type bridge struct{ a, b int }
	bs := make([]bridge, m)
	for i := 0; i < m; i++ {
		a, b := readInt(), readInt()
		add(a, b)
		add(b, a)
		bs[i] = bridge{a, b}
	}

	ans := 0
	for _, bv := range bs {
		ba, bb = bv.a, bv.b
		seen = make(map[int]bool, n)
		dfs(1)
		if len(seen) != n {
			ans++
		}
	}
	fmt.Println(ans)
}
