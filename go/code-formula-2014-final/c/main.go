package main

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func scanString() string {
	sc.Scan()
	return sc.Text()
}

func init() {
	sc.Split(bufio.ScanLines)
	sc.Buffer(make([]byte, 100001), 100001*100)
}

func main() {
	s := scanString()
	xm := make(map[string]bool)
	for _, ss := range strings.Split(s, " ") {
		for i, sss := range strings.Split(ss, "@") {
			if i >= 1 {
				xm[sss] = true
			}
		}
	}
	delete(xm, "")

	xs := make([]string, 0, len(xm))
	for k := range xm {
		xs = append(xs, k)
	}
	sort.Strings(xs)
	wr := bufio.NewWriter(os.Stdout)
	for _, v := range xs {
		wr.WriteString(v + "\n")
	}
	wr.Flush()
}
