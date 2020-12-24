package main

import (
	"fmt"
	"sort"
)

func main() {
	a := make([]int, 3)
	fmt.Scan(&a[0], &a[1], &a[2])
	sort.Ints(a)

	ans := 0
	d1 := a[2] - a[1]
	ans += d1

	d2 := a[2] - (a[0] + d1)
	ans += d2 / 2
	if d2%2 != 0 {
		ans += 2
	}

	fmt.Println(ans)
}
