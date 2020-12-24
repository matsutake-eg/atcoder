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

func main() {
	n, m := scanInt(), scanInt()
	initGraph(n)
	for i := 0; i < m; i++ {
		a, b := scanInt(), scanInt()
		add(a, b)
		add(b, a)
	}

	for i := 1; i <= n; i++ {
		ans := make(map[int]bool)
		for to := range graph[i] {
			for toto := range graph[to] {
				if toto != i && !graph[i][toto] {
					ans[toto] = true
				}
			}
		}
		fmt.Println(len(ans))
	}
}
