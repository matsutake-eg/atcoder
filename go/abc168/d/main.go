package main

import (
	"bufio"
	"container/list"
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

	seen := make([]int, n+1)
	for i := range seen {
		seen[i] = -1
	}
	ls := list.New()
	ls.PushBack(1)
	for ls.Len() > 0 {
		v := ls.Remove(ls.Front()).(int)
		for t := range graph[v] {
			if seen[t] > 0 {
				continue
			}
			seen[t] = v
			ls.PushBack(t)
		}
	}

	exist := true
	for _, v := range seen[2:] {
		if v == -1 {
			exist = false
			break
		}
	}
	var wr = bufio.NewWriter(os.Stdout)
	if exist {
		wr.WriteString("Yes\n")
		for _, v := range seen[2:] {
			wr.WriteString(strconv.Itoa(v) + "\n")
		}
	} else {
		wr.WriteString("No\n")
	}
	wr.Flush()
}
