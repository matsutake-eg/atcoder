package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var pq []int

func push(x int) {
	p := sort.Search(len(pq), func(i int) bool { return pq[i] <= x })
	if p == len(pq) {
		pq = append(pq, x)
		return
	}
	pq = append(pq[:p+1], pq[p:]...)
	pq[p] = x
}

func pop() int {
	if len(pq) == 0 {
		return 0
	}

	ans := pq[0]
	pq = pq[1:]
	return ans
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	ts := make(map[int][]int)
	for i := 0; i < n; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		b, _ := strconv.Atoi(sc.Text())
		ts[a] = append(ts[a], b)
	}

	pq = make([]int, 0, n)
	ans := 0
	for i := 1; i <= m; i++ {
		if _, ok := ts[i]; ok {
			for _, v := range ts[i] {
				push(v)
			}
		}
		ans += pop()
	}
	fmt.Println(ans)
}
