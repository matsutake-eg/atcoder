package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, x, y int
	fmt.Scan(&n, &x, &y)

	type edge struct{ v, d int }
	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		seen := make([]int, n+1)
		seen[i] = 1
		ls := list.New()
		ls.PushBack(i)
		for ls.Len() > 0 {
			v := ls.Remove(ls.Front()).(int)
			if nv := v - 1; nv >= 1 && seen[nv] == 0 {
				seen[nv] = seen[v] + 1
				ans[seen[nv]-1]++
				ls.PushBack(nv)
			}
			if nv := v + 1; nv <= n && seen[nv] == 0 {
				seen[nv] = seen[v] + 1
				ans[seen[nv]-1]++
				ls.PushBack(nv)
			}
			if v == x && seen[y] == 0 {
				seen[y] = seen[v] + 1
				ans[seen[y]-1]++
				ls.PushBack(y)
			}
			if v == y && seen[x] == 0 {
				seen[x] = seen[v] + 1
				ans[seen[x]-1]++
				ls.PushBack(x)
			}
		}
	}

	var wr = bufio.NewWriter(os.Stdout)
	for i := 1; i < n; i++ {
		wr.WriteString(strconv.Itoa(ans[i]/2) + "\n")
	}
	wr.Flush()
}
