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
	n, k := scanInt(), scanInt()
	r := make([]int, n)
	for i := range r {
		r[i] = scanInt()
	}

	sort.Sort(sort.Reverse(sort.IntSlice(r)))
	ans := 0.0
	for i := k - 1; i >= 0; i-- {
		ans = (ans + float64(r[i])) / 2
	}
	fmt.Println(ans)
}
