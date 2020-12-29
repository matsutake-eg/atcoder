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
	n, m := scanInt(), scanInt()
	as := make([]int, n)
	sum := 0
	for i := range as {
		as[i] = scanInt()
		sum += as[i]
	}
	sort.Ints(as)

	type card struct{ n, s int }
	cs := make([]card, m)
	for i := range cs {
		cs[i].s, cs[i].n = scanInt(), scanInt()
	}
	sort.Slice(cs, func(i, j int) bool { return cs[i].n > cs[j].n })

	p := 0
lb:
	for _, v := range cs {
		for i := 0; i < v.s; i++ {
			if p >= len(as) || v.n < as[p] {
				break lb
			}
			sum += v.n - as[p]
			p++
		}
	}
	fmt.Println(sum)
}
