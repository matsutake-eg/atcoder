package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t string
	fmt.Scan(&t)

	wr := bufio.NewWriter(os.Stdout)
	for _, r := range t {
		if r == '?' {
			wr.WriteString("D")
		} else {
			wr.WriteString(string(r))
		}
	}
	wr.WriteString("\n")
	wr.Flush()
}
