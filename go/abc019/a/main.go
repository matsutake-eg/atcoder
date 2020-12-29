package main

import (
	"fmt"
	"sort"
)

func main() {
	xs := make([]int, 3)
	fmt.Scan(&xs[0], &xs[1], &xs[2])

	sort.Ints(xs)
	fmt.Println(xs[1])
}
