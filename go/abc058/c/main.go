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
}

func main() {
	n := scanInt()
	xm := make(map[rune]int)
	s1 := scanString()
	for _, r := range s1 {
		xm[r]++
	}
	for i := 0; i < n-1; i++ {
		s := scanString()
		tm := make(map[rune]int)
		for _, r := range s {
			if _, ok := xm[r]; ok && tm[r] < xm[r] {
				tm[r]++
			}
		}
		xm = tm
	}

	var wr = bufio.NewWriter(os.Stdout)
	for _, r := range "abcdefghijklmnopqrstuvwxyz" {
		if cnt, ok := xm[r]; ok {
			wr.WriteString(strings.Repeat(string(r), cnt))
		}
	}
	wr.WriteString("\n")
	wr.Flush()
}
