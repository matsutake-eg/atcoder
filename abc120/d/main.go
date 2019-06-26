package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	uft      []int
	children map[int]int64
)

func find(x int) int {
	if uft[x] == x {
		return x
	}

	uft[x] = find(uft[x])
	return uft[x]
}

func union(a, b int) {
	uft[a] = uft[b]
	children[b] += children[a]
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
	children = make(map[int]int64)
	for i := range uft {
		uft[i] = i
		children[i]++
	}

	ans := make([]int64, m)
	ans[m-1] = int64(n * (n - 1) / 2)
	for i := m - 1; i >= 1; i-- {
		ar := find(as[i])
		br := find(bs[i])

		if ar != br {
			ans[i-1] = ans[i] - children[ar]*children[br]
			union(ar, br)
			continue
		}
		ans[i-1] = ans[i]
	}

	for _, v := range ans {
		fmt.Println(v)
	}
}
