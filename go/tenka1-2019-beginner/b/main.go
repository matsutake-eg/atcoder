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

	wr := bufio.NewWriter(os.Stdout)
	for _, r := range s {
		if r != rune(s[k-1]) {
			wr.WriteString("*")
		} else {
			wr.WriteString(string(r))
		}
	}
	wr.WriteString("\n")
	wr.Flush()
}
