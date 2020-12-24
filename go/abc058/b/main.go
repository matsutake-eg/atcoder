package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var o, e string
	fmt.Scan(&o, &e)

	var wr = bufio.NewWriter(os.Stdout)
	for i := 0; i < len(o); i++ {
		if i < len(e) {
			wr.WriteString(o[i : i+1])
			wr.WriteString(e[i : i+1])
		} else {
			wr.WriteString(o[i : i+1])
		}
	}
	wr.WriteString("\n")
	wr.Flush()
}
