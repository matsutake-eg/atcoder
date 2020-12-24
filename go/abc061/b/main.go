package main

import (
	"bufio"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	iv, _ := strconv.Atoi(sc.Text())
	return iv
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	n, m := scanInt(), scanInt()
	ans := make([]int, n+1)
	for i := 0; i < m; i++ {
		a, b := scanInt(), scanInt()
		ans[a]++
		ans[b]++
	}

	var wr = bufio.NewWriter(os.Stdout)
	for _, v := range ans[1:] {
		wr.WriteString(strconv.Itoa(v) + "\n")
	}
	wr.Flush()
}
