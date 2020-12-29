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

	n %= 30
	ans := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < n; i++ {
		from, to := i%5, i%5+1
		ans[from], ans[to] = ans[to], ans[from]
	}

	var wr = bufio.NewWriter(os.Stdout)
	for _, v := range ans {
		wr.WriteString(strconv.Itoa(v))
	}
	wr.WriteString("\n")
	wr.Flush()
}
