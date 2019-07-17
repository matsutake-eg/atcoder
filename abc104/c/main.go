package main

import "fmt"

var (
	d, g int
	p, c []int
	cpt  []bool
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func dfs(i int) int {
	if i < len(cpt) {
		cpt[i] = true
		x := dfs(i + 1)
		cpt[i] = false
		y := dfs(i + 1)
		return min(x, y)
	}

	cnt, total := 0, 0
	for i := 0; i < d; i++ {
		if cpt[i] {
			total += (i+1)*100*p[i] + c[i]
			cnt += p[i]
		}
	}
	if total >= g {
		return cnt
	}

	maxB, maxP := 0, 0
	for i := 0; i < d; i++ {
		if !cpt[i] {
			if b := (i + 1) * 100; b > maxB {
				maxB = b
				maxP = p[i]
			}
		}
	}
	n := (g - total + maxB - 1) / maxB
	if n >= maxP {
		return 10000
	}
	return cnt + n
}

func main() {
	fmt.Scan(&d, &g)

	p = make([]int, d)
	c = make([]int, d)
	for i := 0; i < d; i++ {
		fmt.Scan(&p[i], &c[i])
	}

	cpt = make([]bool, d)
	fmt.Println(dfs(0))
}
