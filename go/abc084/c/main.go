package main

import (
	"bufio"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	iv, _ := strconv.Atoi(sc.Text())
	return iv
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	n := readInt()
	type station struct{ c, s, f int }
	st := make([]station, n-1)
	for i := range st {
		c, s, f := readInt(), readInt(), readInt()
		st[i] = station{c, s, f}
	}

	var wr = bufio.NewWriter(os.Stdout)
	for i := range st {
		ans := 0
		for _, s := range st[i:] {
			if ans <= s.s {
				ans = s.s
			} else {
				d := (ans - s.s + s.f - 1) / s.f
				ans = d*s.f + s.s
			}
			ans += s.c
		}
		wr.WriteString(strconv.Itoa(ans) + "\n")
	}
	wr.WriteString(strconv.Itoa(0) + "\n")
	wr.Flush()
}
