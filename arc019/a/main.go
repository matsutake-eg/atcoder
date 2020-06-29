package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s string
	fmt.Scan(&s)

	wr := bufio.NewWriter(os.Stdout)
	for _, r := range s {
		switch r {
		case 'O', 'D':
			wr.WriteString("0")
		case 'I':
			wr.WriteString("1")
		case 'Z':
			wr.WriteString("2")
		case 'S':
			wr.WriteString("5")
		case 'B':
			wr.WriteString("8")
		default:
			wr.WriteString(string(r))
		}
	}
	wr.WriteString("\n")
	wr.Flush()
}
