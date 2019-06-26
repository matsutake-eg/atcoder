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
	ch   []int
	ans  []int64
)

func find(x int) int {
	if uft[x] == x {
		return x
	}

	uft[x] = find(uft[x])
	return uft[x]
}

func union(a, b, i int) {
	ans[i-1] = ans[i]
	ar, br := find(a), find(b)
	if ar == br {
		return
	}

	ans[i-1] -= int64(ch[ar] * ch[br])

	if rank[ar] < rank[br] {
		uft[ar] = br
		ch[br] += ch[ar]
		return
	} else if rank[ar] == rank[br] {
		rank[ar]++
	}
	uft[br] = ar
	ch[ar] += ch[br]
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
	ch = make([]int, n+1)
	for i := range uft {
		uft[i] = i
		ch[i] = 1
	}

	ans = make([]int64, m)
	ans[m-1] = int64(n * (n - 1) / 2)
	for i := m - 1; i >= 1; i-- {
		union(as[i], bs[i], i)
	}

	for _, v := range ans {
		fmt.Println(v)
	}
}
