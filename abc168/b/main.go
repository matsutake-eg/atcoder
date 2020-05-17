package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		k int
		s string
	)
	fmt.Scan(&k, &s)

	var wr = bufio.NewWriter(os.Stdout)
	for i := range s {
		if i > k-1 {
			wr.WriteString("...")
			break
		}
		wr.WriteString(string(s[i]))
	}
	wr.WriteString("\n")
	wr.Flush()
}
