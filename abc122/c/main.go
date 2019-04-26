package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, q int
	var s string
	fmt.Scan(&n, &q, &s)

	xs := make([]int, n)
	xs[0] = 0
	for i := 1; i < n; i++ {
		if s[i-1] == 'A' && s[i] == 'C' {
			xs[i] = xs[i-1] + 1
		} else {
			xs[i] = xs[i-1]
		}
	}

	var l, r int
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	for i := 0; i < q; i++ {
		sc.Scan()
		l, _ = strconv.Atoi(sc.Text())
		sc.Scan()
		r, _ = strconv.Atoi(sc.Text())
		fmt.Println(xs[r-1] - xs[l-1])
	}
}
