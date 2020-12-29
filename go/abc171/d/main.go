package main

import (
	"bufio"
	"os"
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
	xm := make(map[int]int)
	sum := 0
	for i := 0; i < n; i++ {
		a := scanInt()
		xm[a]++
		sum += a
	}

	wr := bufio.NewWriter(os.Stdout)
	q := scanInt()
	for i := 0; i < q; i++ {
		b, c := scanInt(), scanInt()
		sum -= b * xm[b]
		sum += c * xm[b]
		wr.WriteString(strconv.Itoa(sum) + "\n")
		xm[c] += xm[b]
		delete(xm, b)
	}
	wr.Flush()
}
