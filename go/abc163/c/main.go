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
	xs := make([]int, n+1)
	for i := 1; i < n; i++ {
		a := scanInt()
		xs[a]++
	}

	var wr = bufio.NewWriter(os.Stdout)
	for i := 1; i <= n; i++ {
		wr.WriteString(strconv.Itoa(xs[i]) + "\n")
	}
	wr.Flush()
}
