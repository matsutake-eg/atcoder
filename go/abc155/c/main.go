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
	xm := make(map[string]int, n)
	for i := 0; i < n; i++ {
		s := scanString()
		xm[s]++
	}

	type poll struct {
		s string
		c int
	}
	ps := make([]poll, 0, len(xm))
	for s, c := range xm {
		ps = append(ps, poll{s, c})
	}
	sort.Slice(ps, func(i, j int) bool { return ps[i].c > ps[j].c })

	ans := make([]string, 0, len(xm))
	pmx := ps[0]
	ans = append(ans, pmx.s)
	for _, p := range ps[1:] {
		if p.c != pmx.c {
			break
		}
		ans = append(ans, p.s)
	}
	sort.Strings(ans)

	var wr = bufio.NewWriter(os.Stdout)
	for _, s := range ans {
		wr.WriteString(s + "\n")
	}
	wr.Flush()
}
