package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var a, b, k int
	fmt.Scan(&a, &b, &k)

	wr := bufio.NewWriter(os.Stdout)
	if a+k > b-k {
		for i := a; i <= b; i++ {
			wr.WriteString(strconv.Itoa(i) + "\n")
		}
		wr.Flush()
		return
	}

	for i := a; i < a+k; i++ {
		wr.WriteString(strconv.Itoa(i) + "\n")
	}
	for i := b - k + 1; i <= b; i++ {
		wr.WriteString(strconv.Itoa(i) + "\n")
	}
	wr.Flush()
}
