package main

import (
	"fmt"
	"sort"
)

func main() {
	a := make([]int, 3)
	fmt.Scan(&a[0], &a[1], &a[2])

	sort.Ints(a)
	fmt.Println(a[2] - a[1] + a[1] - a[0])
}
