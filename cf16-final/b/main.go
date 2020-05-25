package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	sum := 0
	wr := bufio.NewWriter(os.Stdout)
	for i := 1; ; i++ {
		sum += i
		if d := sum - n; d > 0 {
			for j := 1; j <= i; j++ {
				if j == d {
					continue
				}
				wr.WriteString(strconv.Itoa(j) + "\n")
			}
			wr.Flush()
			return
		}
	}
}
