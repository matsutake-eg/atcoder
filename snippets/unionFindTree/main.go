package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	uft  []int
	rank []int
)

func find(x int) int {
	if uft[x] == x {
		return x
	}

	uft[x] = find(uft[x])
	return uft[x]
}

func union(a, b int) {
	ar, br := find(a), find(b)
	if ar == br {
		return
	}

	if rank[ar] < rank[br] {
		uft[ar] = br
		return
	} else if rank[ar] == rank[br] {
		rank[ar]++
	}
	uft[br] = ar
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	as := make([]int, m)
	bs := make([]int, m)
	for i := 0; i < m; i++ {
		sc.Scan()
		as[i], _ = strconv.Atoi(sc.Text())
		sc.Scan()
		bs[i], _ = strconv.Atoi(sc.Text())
	}

	uft = make([]int, n+1)
	rank = make([]int, n+1)
	for i := range uft {
		uft[i] = i
	}

	for i := m - 1; i >= 1; i-- {
		union(as[i], bs[i])
	}
	fmt.Println(uft)
}
