package main

import (
	"bufio"
	"fmt"
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

func main() {
	n, m, x, y := scanInt(), scanInt(), scanInt(), scanInt()
	a := make([]int, n)
	for i := range a {
		a[i] = scanInt()
	}
	b := make([]int, m)
	for i := range b {
		b[i] = scanInt()
	}

	ans := 0
	ap, bp := 0, 0
	for {
		if ap > a[n-1] {
			break
		}
		ai := sort.SearchInts(a, ap)
		bp = a[ai] + x

		if bp > b[m-1] {
			break
		}
		bi := sort.SearchInts(b, bp)
		ap = b[bi] + y
		ans++
	}
	fmt.Println(ans)
}
