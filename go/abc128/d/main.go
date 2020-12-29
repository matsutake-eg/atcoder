package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	vs := make([]int, n)
	for i := range vs {
		fmt.Scan(&vs[i])
	}
	min := k
	if n < min {
		min = n
	}

	ans := 0
	for a := 0; a <= min; a++ {
		for b := 0; b <= min-a; b++ {
			xs := make([]int, 0, a+b)
			xs = append(xs, vs[0:a]...)
			xs = append(xs, vs[len(vs)-b:len(vs)]...)
			sort.Ints(xs)

			sum := 0
			d := k - (a + b)
			for _, v := range xs {
				if v < 0 && d > 0 {
					d--
					continue
				}
				sum += v
			}
			if sum > ans {
				ans = sum
			}
		}
	}
	fmt.Println(ans)
}
