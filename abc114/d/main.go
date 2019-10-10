package main

import "fmt"

var pfm = make(map[int]int)

func prmFac(x int) map[int]int {
	for i := 2; i*i <= x; i++ {
		for x%i == 0 {
			x /= i
			pfm[i]++
		}
	}
	if x > 1 {
		pfm[x]++
	}
	return pfm
}

func p(val int) int {
	pc := 0
	for _, v := range pfm {
		if v >= val-1 {
			pc++
		}
	}
	return pc
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		prmFac(i)
	}
	fmt.Println(p(75) + p(25)*(p(3)-1) + p(15)*(p(5)-1) + p(5)*(p(5)-1)*(p(3)-2)/2)
}
