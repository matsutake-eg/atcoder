package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		s string
		w int
	)
	fmt.Scan(&s, &w)

	var wr = bufio.NewWriter(os.Stdout)
	for i := 0; i < len(s); i += w {
		wr.WriteString(string(s[i]))
	}
	wr.WriteString("\n")
	wr.Flush()
}
