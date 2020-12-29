package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const max = 100000000000

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var a, b, q int
	fmt.Scan(&a, &b, &q)

	sc := bufio.NewScanner(os.Stdin)
	s := make([]int, a+2)
	s[0], s[len(s)-1] = -max, max
	for i := 1; i < len(s)-1; i++ {
		sc.Scan()
		s[i], _ = strconv.Atoi(sc.Text())
	}
	t := make([]int, b+2)
	t[0], t[len(t)-1] = -max, max
	for i := 1; i < len(t)-1; i++ {
		sc.Scan()
		t[i], _ = strconv.Atoi(sc.Text())
	}
	x := make([]int, q)
	for i := range x {
		sc.Scan()
		x[i], _ = strconv.Atoi(sc.Text())
	}

	wr := bufio.NewWriter(os.Stdout)
	for _, vx := range x {
		ss := sort.SearchInts(s, vx)
		st := sort.SearchInts(t, vx)
		ans := max * 2
		for _, vs := range s[ss-1 : ss+1] {
			for _, vt := range t[st-1 : st+1] {
				if d := abs(vx-vs) + abs(vs-vt); d < ans {
					ans = d
				}
				if d := abs(vx-vt) + abs(vs-vt); d < ans {
					ans = d
				}
			}
		}
		wr.WriteString(strconv.Itoa(ans) + "\n")
	}
	wr.Flush()
}
