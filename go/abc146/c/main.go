package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var a, b, x int
	fmt.Scan(&a, &b, &x)

	if a+b > x {
		fmt.Println(0)
		return
	}

	idx := sort.Search(1000000000, func(i int) bool {
		return x <= a*i+b*len(strconv.Itoa(i))
	})
	if a*idx+b*len(strconv.Itoa(idx)) > x {
		fmt.Println(idx - 1)
		return
	}
	fmt.Println(idx)
}
