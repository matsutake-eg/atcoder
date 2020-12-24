package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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
	sc.Buffer(make([]byte, 100000), 100000000)
}

func main() {
	n := scanInt()
	xs := make([]string, 0, n)
	for i := 0; i < n; i++ {
		w := scanString()
		s := ""
		for _, r := range w {
			ignore := false
			for _, ir := range "aeiouy.," {
				if strings.ToLower(string(r)) == string(ir) {
					ignore = true
					break
				}
			}
			if !ignore {
				s += strings.ToLower(string(r))
			}
		}
		if len(s) > 0 {
			xs = append(xs, s)
		}
	}

	xm := make(map[rune]int)
	xm['b'], xm['c'] = 1, 1
	xm['d'], xm['w'] = 2, 2
	xm['t'], xm['j'] = 3, 3
	xm['f'], xm['q'] = 4, 4
	xm['l'], xm['v'] = 5, 5
	xm['s'], xm['x'] = 6, 6
	xm['p'], xm['m'] = 7, 7
	xm['h'], xm['k'] = 8, 8
	xm['n'], xm['g'] = 9, 9
	xm['z'], xm['r'] = 0, 0
	wr := bufio.NewWriter(os.Stdout)
	for i, s := range xs {
		for _, r := range s {
			wr.WriteString(strconv.Itoa(xm[r]))
		}
		if i < len(xs)-1 {
			wr.WriteString(" ")
		}
	}
	wr.WriteString("\n")
	wr.Flush()
}
