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

func main() {
	n := scanInt()
	type student struct{ n, a int }
	sts := make([]student, n)
	for i := range sts {
		sts[i].n = i + 1
		sts[i].a = scanInt()
	}
	sort.Slice(sts, func(i, j int) bool { return sts[i].a < sts[j].a })

	wr := bufio.NewWriter(os.Stdout)
	for i, v := range sts {
		wr.WriteString(strconv.Itoa(v.n))
		if i < len(sts)-1 {
			wr.WriteString(" ")
		} else {
			wr.WriteString("\n")
		}
	}
	wr.Flush()
}
