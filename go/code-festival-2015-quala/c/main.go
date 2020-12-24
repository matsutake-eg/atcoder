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
	n, t := scanInt(), scanInt()
	type w struct{ a, b, d int }
	ws := make([]w, n)
	asum, bsum := 0, 0
	for i := 0; i < n; i++ {
		a, b := scanInt(), scanInt()
		asum += a
		bsum += b
		ws[i] = w{a, b, a - b}
	}
	sort.Slice(ws, func(i, j int) bool { return ws[i].d > ws[j].d })

	if asum <= t {
		fmt.Println(0)
		return
	}
	for i := range ws {
		asum -= ws[i].d
		if asum <= t {
			fmt.Println(i + 1)
			return
		}
	}
	fmt.Println(-1)
}
