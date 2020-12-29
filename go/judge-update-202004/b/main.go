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
	type ball struct{ x, c int }
	bs := make([]ball, n)
	for i := range bs {
		x, c := scanInt(), scanString()
		bs[i].x = x
		if c == "R" {
			bs[i].c = 1
		} else {
			bs[i].c = 2
		}
	}

	sort.Slice(bs, func(i, j int) bool {
		if bs[i].c == bs[j].c {
			return bs[i].x < bs[j].x
		}
		return bs[i].c < bs[j].c
	})
	var wr = bufio.NewWriter(os.Stdout)
	for _, b := range bs {
		wr.WriteString(strconv.Itoa(b.x) + "\n")
	}
	wr.Flush()
}
