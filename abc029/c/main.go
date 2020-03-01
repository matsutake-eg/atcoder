package main

import (
	"bufio"
	"fmt"
	"os"
)

var az string = "abc"
var wr = bufio.NewWriter(os.Stdout)

func rec(x int, s string) {
	if x == 0 {
		wr.WriteString(s + "\n")
		return
	}

	for _, r := range az {
		rec(x-1, s+string(r))
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	rec(n, "")
	wr.Flush()
}
