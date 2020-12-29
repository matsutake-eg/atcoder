package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	v := make([]float64, n)
	for i := range v {
		fmt.Scan(&v[i])
	}

	sort.Float64s(v)
	ans := v[0]
	for _, x := range v[1:] {
		ans = (ans + x) / 2
	}
	fmt.Println(ans)
}
