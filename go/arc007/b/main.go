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
	n, m := scanInt(), scanInt()
	xs := make([]int, n+1)
	for i := range xs {
		xs[i] = i
	}

	for i := 0; i < m; i++ {
		d := scanInt()
		for j := range xs {
			if xs[j] == d {
				xs[j] = xs[0]
				xs[0] = d
				break
			}
		}
	}

	wr := bufio.NewWriter(os.Stdout)
	for _, v := range xs[1:] {
		wr.WriteString(strconv.Itoa(v) + "\n")
	}
	wr.Flush()
}
