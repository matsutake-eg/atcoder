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
	n := scanInt()
	xm := make(map[int]bool)
	as := make([]int, 0, n)
	for i := 0; i < n; i++ {
		a := scanInt()
		if !xm[a] {
			xm[a] = true
			as = append(as, a)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(as)))
	fmt.Println(as[1])
}
