package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func main() {
	n := 1000
	ans := make([]string, n)
	for i := range ans {
		ans[i] = strconv.Itoa(i + 1)
	}
	sort.Strings(ans)

	wr := bufio.NewWriter(os.Stdout)
	for _, v := range ans {
		wr.WriteString(v + "\n")
	}
	wr.Flush()
}
