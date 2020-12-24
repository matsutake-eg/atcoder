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
	n, m := scanInt(), scanInt()
	type shop struct{ p, s int }
	s := make([]shop, n)
	for i := 0; i < n; i++ {
		s[i].p, s[i].s = scanInt(), scanInt()
	}
	sort.Slice(s, func(i, j int) bool { return s[i].p < s[j].p })

	count := 0
	sum := 0
	for i := 0; i < n; i++ {
		if v := count + s[i].s; v >= m {
			sum += s[i].p * (m - count)
			break
		} else {
			count = v
			sum += s[i].p * s[i].s
		}
	}
	fmt.Println(sum)
}
