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

var t [][]int

func dfs(n, x int) bool {
	if n >= len(t) {
		return x == 0
	}

	for _, tv := range t[n] {
		if dfs(n+1, tv^x) {
			return true
		}
	}
	return false
}

func main() {
	n, k := scanInt(), scanInt()
	t = make([][]int, n)
	for i := range t {
		t[i] = make([]int, k)
		for j := range t[i] {
			t[i][j] = scanInt()
		}
	}

	if dfs(0, 0) {
		fmt.Println("Found")
	} else {
		fmt.Println("Nothing")
	}
}
