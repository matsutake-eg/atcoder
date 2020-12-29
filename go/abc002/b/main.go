package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var w string
	fmt.Scan(&w)

	var wr = bufio.NewWriter(os.Stdout)
	for _, r := range w {
		switch r {
		case 'a':
			continue
		case 'i':
			continue
		case 'u':
			continue
		case 'e':
			continue
		case 'o':
			continue
		}
		wr.WriteString(string(r))
	}
	wr.WriteString("\n")
	wr.Flush()
}
