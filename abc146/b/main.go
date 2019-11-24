package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		n int
		s string
	)
	fmt.Scan(&n, &s)

	a := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	m := make(map[rune]int)
	for i, r := range a {
		m[r] = i
	}

	var wr = bufio.NewWriter(os.Stdout)
	for _, r := range s {
		wr.WriteString(string(a[(m[r]+n)%26]))
	}
	wr.WriteString("\n")
	wr.Flush()
}
