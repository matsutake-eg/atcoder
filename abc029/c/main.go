package main

import (
	"bufio"
	"fmt"
	"os"
)

var wr = bufio.NewWriter(os.Stdout)

func dfs(i int, s string) {
	if i == 0 {
		wr.WriteString(s + "\n")
		return
	}
	for _, r := range "abc" {
		dfs(i-1, s+string(r))
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	dfs(n, "")
	wr.Flush()
}
