package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	alice, bob := 0, 0
	for i := range a {
		if i%2 == 0 {
			alice += a[i]
			continue
		}
		bob += a[i]
	}
	fmt.Println(alice - bob)
}
