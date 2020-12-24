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
	b := make([]int, 10)
	for i := range b {
		b[i] = scanInt()
	}

	n := scanInt()
	type no struct{ a, x int }
	a := make([]no, n)
	for i := range a {
		a[i].a = scanInt()
		sa := strconv.Itoa(a[i].a)
		d := 1
		for j := len(sa) - 1; j >= 0; j-- {
			for k := 0; k < len(b); k++ {
				if string(sa[j]) == strconv.Itoa(b[k]) {
					a[i].x += k * d
					break
				}
			}
			d *= 10
		}
	}
	sort.Slice(a, func(i, j int) bool { return a[i].x < a[j].x })

	wr := bufio.NewWriter(os.Stdout)
	for _, v := range a {
		wr.WriteString(strconv.Itoa(v.a) + "\n")
	}
	wr.Flush()
}
