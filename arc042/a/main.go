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
	a := make([]int, m)
	for i := range a {
		a[i] = scanInt()
	}

	xs := make([]bool, n+1)
	wr := bufio.NewWriter(os.Stdout)
	for i := m - 1; i >= 0; i-- {
		if xs[a[i]] {
			continue
		}
		wr.WriteString(strconv.Itoa(a[i]) + "\n")
		xs[a[i]] = true
	}
	for i := 1; i <= n; i++ {
		if !xs[i] {
			wr.WriteString(strconv.Itoa(i) + "\n")
		}
	}
	wr.Flush()
}
