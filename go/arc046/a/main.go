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

	x := (n-1)/9 + 1
	r := (n-1)%9 + 1
	wr := bufio.NewWriter(os.Stdout)
	for i := 0; i < x; i++ {
		wr.WriteString(strconv.Itoa(r))
	}
	wr.WriteString("\n")
	wr.Flush()
}
