package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var h, w, a, b int
	fmt.Scan(&h, &w, &a, &b)

	wr := bufio.NewWriter(os.Stdout)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if (i+1 <= b && j+1 <= a) || (i+1 > b && j+1 > a) {
				wr.WriteString("1")
				continue
			}
			wr.WriteString("0")
		}
		wr.WriteString("\n")
	}
	wr.Flush()
}
