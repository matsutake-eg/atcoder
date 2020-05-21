package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		n, a, b int
		s       string
	)
	fmt.Scan(&n, &a, &b, &s)

	x, y := 0, 0
	var wr = bufio.NewWriter(os.Stdout)
	for _, v := range s {
		switch v {
		case 'a':
			if x+y < a+b {
				x++
				wr.WriteString("Yes")
			} else {
				wr.WriteString("No")
			}
		case 'b':
			if x+y < a+b && y < b {
				y++
				wr.WriteString("Yes")
			} else {
				wr.WriteString("No")
			}
		case 'c':
			wr.WriteString("No")
		}
		wr.WriteString("\n")
	}
	wr.Flush()
}
