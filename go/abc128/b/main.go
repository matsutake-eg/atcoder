package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	iv, _ := strconv.Atoi(scanString())
	return iv
}

func scanString() string {
	sc.Scan()
	return sc.Text()
}

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 100001), 100001*100)
}

func main() {
	n := scanInt()
	type restaurant struct {
		n    string
		p, i int
	}
	rs := make([]restaurant, n)
	for i := 0; i < n; i++ {
		rs[i].i = i + 1
		rs[i].n, rs[i].p = scanString(), scanInt()
	}

	sort.Slice(rs, func(i, j int) bool {
		if rs[i].n == rs[j].n {
			return rs[i].p > rs[j].p
		}
		return rs[i].n < rs[j].n
	})

	var wr = bufio.NewWriter(os.Stdout)
	for _, v := range rs {
		wr.WriteString(strconv.Itoa(v.i) + "\n")
	}
	wr.Flush()
}
