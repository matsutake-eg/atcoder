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
	xs := make([]int, 0, n)
	for i := 0; i < n; i++ {
		s := scanInt()
		if s != 0 {
			xs = append(xs, s)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(xs)))

	wr := bufio.NewWriter(os.Stdout)
	q := scanInt()
	for i := 0; i < q; i++ {
		k := scanInt()
		if k >= len(xs) {
			wr.WriteString("0")
		} else {
			wr.WriteString(strconv.Itoa(xs[k] + 1))
		}
		wr.WriteString("\n")
	}
	wr.Flush()
}
