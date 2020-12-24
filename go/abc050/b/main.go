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
	t := make([]int, n)
	sum := 0
	for i := range t {
		t[i] = scanInt()
		sum += t[i]
	}
	m := scanInt()
	var wr = bufio.NewWriter(os.Stdout)
	for i := 0; i < m; i++ {
		p, x := scanInt(), scanInt()
		wr.WriteString(strconv.Itoa(sum-t[p-1]+x) + "\n")
	}
	wr.Flush()
}
