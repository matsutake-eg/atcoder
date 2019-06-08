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
	var sum int
	for i := range s[:len(xs)-1] {
		if s[i] == 'A' && s[i+1] == 'C' {
			sum++
		}
		xs[i+1] = sum
	}

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	var l, r int
	for i := 0; i < q; i++ {
		sc.Scan()
		l, _ = strconv.Atoi(sc.Text())
		sc.Scan()
		r, _ = strconv.Atoi(sc.Text())
		fmt.Println(xs[r-1] - xs[l-1])
	}
}
