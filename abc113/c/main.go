package main

import (
	"bufio"
	"fmt"
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
	_, m := scanInt(), scanInt()
	type city struct{ i, p, y int }
	cs := make([]city, m)
	for i := range cs {
		cs[i].i = i
		cs[i].p, cs[i].y = scanInt(), scanInt()
	}
	sort.Slice(cs, func(i, j int) bool {
		if cs[i].p == cs[j].p {
			return cs[i].y < cs[j].y
		}
		return cs[i].p < cs[j].p
	})

	xs := make([]string, m)
	p, c := 0, 0
	for _, v := range cs {
		if p != v.p {
			p = v.p
			c = 0
		}
		c++
		xs[v.i] = fmt.Sprintf("%06d%06d", p, c)
	}

	wr := bufio.NewWriter(os.Stdout)
	for _, v := range xs {
		wr.WriteString(v + "\n")
	}
	wr.Flush()
}
