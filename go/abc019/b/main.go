package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)

	x := rune(s[0])
	cnt := 1
	var wr = bufio.NewWriter(os.Stdout)
	for _, r := range s[1:] {
		if x != r {
			wr.WriteString(string(x) + strconv.Itoa(cnt))
			x = r
			cnt = 1
		} else {
			cnt++
		}
	}
	wr.WriteString(string(x) + strconv.Itoa(cnt))
	wr.WriteString("\n")
	wr.Flush()
}
