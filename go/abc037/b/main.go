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
	n, q := scanInt(), scanInt()
	ans := make([]int, n)
	for i := 0; i < q; i++ {
		l, r, t := scanInt(), scanInt(), scanInt()
		for j := l - 1; j <= r-1; j++ {
			ans[j] = t
		}
	}

	var wr = bufio.NewWriter(os.Stdout)
	for _, v := range ans {
		wr.WriteString(strconv.Itoa(v) + "\n")
	}
	wr.Flush()
}
