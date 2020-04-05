package main

import (
	"bufio"
	"os"
	"sort"
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

func gcd(x, y int) int {
	for y > 0 {
		x, y = y, x%y
	}
	return x
}

func main() {
	n, q := scanInt(), scanInt()
	a := make([]int, n)
	a[0] = scanInt()
	for i := 1; i < n; i++ {
		v := scanInt()
		a[i] = gcd(a[i-1], v)
	}

	var wr = bufio.NewWriter(os.Stdout)
	for i := 0; i < q; i++ {
		s := scanInt()
		idx := sort.Search(n, func(i int) bool { return gcd(s, a[i]) == 1 })
		if idx == n {
			idx--
		}
		if g := gcd(s, a[idx]); g == 1 {
			wr.WriteString(strconv.Itoa(idx+1) + "\n")
		} else {
			wr.WriteString(strconv.Itoa(g) + "\n")
		}
	}
	wr.Flush()
}
