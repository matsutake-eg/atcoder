package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	iv, _ := strconv.Atoi(readStr())
	return iv
}

func readStr() string {
	sc.Scan()
	return sc.Text()
}

type poll struct {
	s string
	c int
}
type polls []poll

func (p polls) Len() int           { return len(p) }
func (p polls) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p polls) Less(i, j int) bool { return p[i].c > p[j].c }

func main() {
	n := readInt()
	m := make(map[string]int)
	for i := 0; i < n; i++ {
		s := readStr()
		m[s]++
	}

	ps := polls(make([]poll, 0, len(m)))
	for s, c := range m {
		ps = append(ps, poll{s, c})
	}
	sort.Sort(ps)

	ans := make([]string, 0, len(m))
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
