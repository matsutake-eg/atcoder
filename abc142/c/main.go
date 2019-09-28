package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type student struct{ no, a int }
type students []student

func (s students) Len() int           { return len(s) }
func (s students) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s students) Less(i, j int) bool { return s[i].a < s[j].a }

func main() {
	var n int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sts := students(make([]student, n))
	for i := range sts {
		sts[i].no = i + 1
		sc.Scan()
		sts[i].a, _ = strconv.Atoi(sc.Text())
	}

	sort.Sort(sts)
	wr := bufio.NewWriter(os.Stdout)
	for i, v := range sts {
		wr.WriteString(strconv.Itoa(v.no))
		if i != len(sts)-1 {
			wr.WriteString(" ")
			continue
		}
		wr.WriteString("\n")
	}
	wr.Flush()
}
