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
	s := make([]string, n)
	for i := range s {
		s[i] = scanString()
	}

	rs := make([]string, n)
	for i := range s {
		r := []rune(s[i])
		for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		rs[i] = string(r)
	}

	sort.Strings(rs)
	wr := bufio.NewWriter(os.Stdout)
	for i := range rs {
		r := []rune(rs[i])
		for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		wr.WriteString(string(r) + "\n")
	}
	wr.Flush()
}
