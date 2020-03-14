package main

import (
	"bufio"
	"fmt"
	"os"
)

var n int
var az string = "abcdefghijklmnopqrstuvwxyz"

var wr = bufio.NewWriter(os.Stdout)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func rec(s string, x int) {
	if len(s) == n {
		wr.WriteString(s + "\n")
		return
	}
	for i := 0; i < x+1; i++ {
		rec(s+az[i:i+1], max(x, i+1))
	}
}

func main() {
	fmt.Scan(&n)

	rec("", 0)
	wr.Flush()
}
