package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	wr := bufio.NewWriter(os.Stdout)

	for i := 1; i <= 127; i++ {
		if i%3 == a && i%5 == b && i%7 == c {
			wr.WriteString(strconv.Itoa(i) + "\n")
		}
	}
	wr.Flush()
}
