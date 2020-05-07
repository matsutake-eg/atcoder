package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		n, k int
		s    string
	)
	fmt.Scan(&n, &s, &k)

	var wr = bufio.NewWriter(os.Stdout)

	for i := 0; i < n; i++ {
		if s[i] != s[k-1] {
			wr.WriteString("*")
		} else {
			wr.WriteString(string(s[i]))
		}
	}
	wr.WriteString("\n")
	wr.Flush()
}
