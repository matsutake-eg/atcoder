package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

type sit struct{ x, i int }
type sits []sit

func (s sits) Len() int           { return len(s) }
func (s sits) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sits) Less(i, j int) bool { return s[i].x < s[j].x }

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
	ss := sits(make([]sit, n))
	for i := range ss {
		ss[i].x, ss[i].i = scanInt(), i
	}

	sort.Sort(ss)
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
