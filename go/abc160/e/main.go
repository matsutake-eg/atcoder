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
	x, y, a, b, c := scanInt(), scanInt(), scanInt(), scanInt(), scanInt()
	p := make([]int, a)
	for i := range p {
		p[i] = scanInt()
	}
	q := make([]int, b)
	for i := range q {
		q[i] = scanInt()
	}
	r := make([]int, c)
	for i := range r {
		r[i] = scanInt()
	}

	sort.Sort(sort.Reverse(sort.IntSlice(p)))
	sort.Sort(sort.Reverse(sort.IntSlice(q)))
	sort.Sort(sort.Reverse(sort.IntSlice(r)))
	ans := 0
	pi, qi, ri := 0, 0, 0
	for cnt := 0; cnt < x+y; cnt++ {
		pv := 0
		if pi < x {
			pv = p[pi]
		}
		qv := 0
		if qi < y {
			qv = q[qi]
		}
		rv := 0
		if ri < c {
			rv = r[ri]
		}

		if pv >= qv && pv >= rv {
			ans += pv
			pi++
		} else if qv >= pv && qv >= rv {
			ans += qv
			qi++
		} else {
			ans += rv
			ri++
		}
	}
	fmt.Println(ans)
}
