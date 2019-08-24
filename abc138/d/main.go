package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, q int
	fmt.Scan(&n, &q)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	eg := make(map[int]map[int]bool)
	for i := 0; i < n-1; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		b, _ := strconv.Atoi(sc.Text())
		if _, ok := eg[a]; !ok {
			eg[a] = make(map[int]bool)
		}
		eg[a][b] = true
		if _, ok := eg[b]; !ok {
			eg[b] = make(map[int]bool)
		}
		eg[b][a] = true
	}

	ans := make([]int, n+1)
	for i := 0; i < q; i++ {
		sc.Scan()
		p, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		x, _ := strconv.Atoi(sc.Text())
		ans[p] += x
	}

	qu := make([]int, 1, n)
	qu[0] = 1
	wk := make([]bool, n+1)
	wk[1] = true
	for {
		p := qu[0]
		for v := range eg[p] {
			if wk[v] {
				continue
			}
			wk[v] = true
			qu = append(qu, v)
			ans[v] += ans[p]
		}
		if len(qu) > 1 {
			qu = qu[1:]
			continue
		}
		break
	}

	wr := bufio.NewWriter(os.Stdout)
	for i := 1; i <= n; i++ {
		wr.WriteString(strconv.Itoa(ans[i]))
		if i != n {
			wr.WriteString(" ")
			continue
		}
		wr.WriteString("\n")
	}
	wr.Flush()
}
