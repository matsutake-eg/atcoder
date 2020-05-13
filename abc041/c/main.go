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
	type student struct{ a, i int }
	sts := make([]student, n)
	for i := range sts {
		sts[i].a, sts[i].i = scanInt(), i+1
	}

	sort.Slice(sts, func(i, j int) bool { return sts[i].a > sts[j].a })
	var wr = bufio.NewWriter(os.Stdout)
	for _, st := range sts {
		wr.WriteString(strconv.Itoa(st.i) + "\n")
	}
	wr.Flush()
}
