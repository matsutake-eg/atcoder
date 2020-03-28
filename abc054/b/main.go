package main

import (
	"bufio"
	"fmt"
	"os"
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
	n, m := scanInt(), scanInt()
	a := make([]string, n)
	for i := range a {
		a[i] = scanString()
	}
	b := make([]string, m)
	for i := range b {
		b[i] = scanString()
	}

	for ai := 0; ai < n-m+1; ai++ {
		for aj := 0; aj < n-m+1; aj++ {
		L:
			for bi := 0; bi < m; bi++ {
				for bj := 0; bj < m; bj++ {
					if a[ai+bi][aj+bj] != b[bi][bj] {
						break L
					}
					if bi == m-1 && bj == m-1 {
						fmt.Println("Yes")
						return
					}
				}
			}
		}
	}
	fmt.Println("No")
}
