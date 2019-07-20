package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type city struct{ index, prefecture, year int }
type cities []city

func (c cities) Len() int      { return len(c) }
func (c cities) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c cities) Less(i, j int) bool {
	if c[i].prefecture == c[j].prefecture {
		return c[i].year < c[j].year
	}
	return c[i].prefecture < c[j].prefecture
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	cs := cities(make([]city, m))
	for i := range cs {
		cs[i].index = i
		sc.Scan()
		cs[i].prefecture, _ = strconv.Atoi(sc.Text())
		sc.Scan()
		cs[i].year, _ = strconv.Atoi(sc.Text())
	}

	sort.Sort(cs)

	xs := make([]string, m)
	p, c := 0, 0
	for _, v := range cs {
		if p != v.prefecture {
			p = v.prefecture
			c = 0
		}
		c++
		xs[v.index] = fmt.Sprintf("%06d%06d", p, c)
	}

	wr := bufio.NewWriter(os.Stdout)
	for _, v := range xs {
		wr.WriteString(v + "\n")
	}
	wr.Flush()
}
