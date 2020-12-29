package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	iv, _ := strconv.Atoi(sc.Text())
	return iv
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	n, h := readInt(), readInt()
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = readInt()
		b[i] = readInt()
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	sort.Sort(sort.Reverse(sort.IntSlice(b)))
	d := 0
	ib := 0
	ans := 0
	for d < h {
		if ib == n || b[ib] <= a[0] {
			ans += (h - d + a[0] - 1) / a[0]
			break
		}
		d += b[ib]
		ans++
		ib++
	}
	fmt.Println(ans)
}
