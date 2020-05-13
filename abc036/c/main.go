package main

import (
	"bufio"
	"os"
	"sort"
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
	n := scanInt()
	type sit struct{ x, i int }
	ss := make([]sit, n)
	for i := range ss {
		ss[i].x, ss[i].i = scanInt(), i
	}

	sort.Slice(ss, func(i, j int) bool { return ss[i].x < ss[j].x })
	cnt := 0
	ans := make([]int, n)
	ans[ss[0].i] = 0
	for i := 1; i < n; i++ {
		if ss[i].x > ss[i-1].x {
			cnt++
		}
		ans[ss[i].i] = cnt
	}

	var wr = bufio.NewWriter(os.Stdout)
	for _, v := range ans {
		wr.WriteString(strconv.Itoa(v) + "\n")
	}
	wr.Flush()
}
