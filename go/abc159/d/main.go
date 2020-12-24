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
	xm := make(map[int]int, n)
	a := make([]int, n)
	for i := range a {
		a[i] = scanInt()
		xm[a[i]]++
	}

	sum := 0
	for _, v := range xm {
		sum += v * (v - 1) / 2
	}

	var wr = bufio.NewWriter(os.Stdout)
	for _, v := range a {
		wr.WriteString(strconv.Itoa(sum-(xm[v]-1)) + "\n")
	}
	wr.Flush()
}
