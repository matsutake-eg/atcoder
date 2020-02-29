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

type student struct{ a, i int }
type students []student

func (s students) Len() int           { return len(s) }
func (s students) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s students) Less(i, j int) bool { return s[i].a > s[j].a }

func main() {
	n := scanInt()
	sts := students(make([]student, n))
	for i := range sts {
		sts[i].a, sts[i].i = scanInt(), i+1
	}

	sort.Sort(sts)
	var wr = bufio.NewWriter(os.Stdout)
	for _, st := range sts {
		wr.WriteString(strconv.Itoa(st.i) + "\n")
	}
	wr.Flush()
}
