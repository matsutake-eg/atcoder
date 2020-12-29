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
	n, q := scanInt(), scanInt()
	xs := make([]bool, n+2)
	for i := 0; i < q; i++ {
		l, r := scanInt(), scanInt()
		xs[l] = !xs[l]
		xs[r+1] = !xs[r+1]
	}

	b := 0
	var wr = bufio.NewWriter(os.Stdout)
	for _, x := range xs[1 : n+1] {
		if x {
			b ^= 1
		}
		wr.WriteString(strconv.Itoa(b))
	}
	wr.WriteString("\n")
	wr.Flush()
}
