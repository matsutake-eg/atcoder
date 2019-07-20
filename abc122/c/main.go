package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		n, q int
		s    string
	)
	fmt.Scan(&n, &q, &s)

	xs := make([]int, n)
	sum := 0
	for i := range s[:len(xs)-1] {
		if s[i] == 'A' && s[i+1] == 'C' {
			sum++
		}
		xs[i+1] = sum
	}

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	var l, r int
	wr := bufio.NewWriter(os.Stdout)
	for i := 0; i < q; i++ {
		sc.Scan()
		l, _ = strconv.Atoi(sc.Text())
		sc.Scan()
		r, _ = strconv.Atoi(sc.Text())
		wr.WriteString(strconv.Itoa(xs[r-1]-xs[l-1]) + "\n")
	}
	wr.Flush()
}
