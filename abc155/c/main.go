package main

import (
	"bufio"
	"container/heap"
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
type prQue []*poll

func (p prQue) Len() int            { return len(p) }
func (p prQue) Less(i, j int) bool  { return p[i].c > p[j].c }
func (p prQue) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p *prQue) Push(x interface{}) { *p = append(*p, x.(*poll)) }
func (p *prQue) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}
func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	n := readInt()
	m := make(map[string]int)
	for i := 0; i < n; i++ {
		s := readStr()
		m[s]++
	}

	pq := make(prQue, 0, len(m))
	for s, c := range m {
		heap.Push(&pq, &poll{s, c})
	}

	ans := make([]string, 0, len(m))
	pm := heap.Pop(&pq).(*poll)
	ans = append(ans, pm.s)
	for len(pq) > 0 {
		p := heap.Pop(&pq).(*poll)
		if p.c == pm.c {
			ans = append(ans, p.s)
		} else {
			break
		}
	}
	sort.Strings(ans)

	var wr = bufio.NewWriter(os.Stdout)
	for _, s := range ans {
		wr.WriteString(s + "\n")
	}
	wr.Flush()
}
