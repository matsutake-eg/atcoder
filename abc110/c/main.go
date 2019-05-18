package main

import (
	"fmt"
	"sort"
)

func getArray(m map[rune]int) (a []int) {
	a = make([]int, len(m))
	index := 0
	for _, v := range m {
		a[index] = v
		index++
	}
	return
}

func main() {
	var s, t string
	fmt.Scan(&s, &t)
	sm := make(map[rune]int)
	for _, v := range s {
		sm[v]++
	}
	sa := getArray(sm)
	sort.Ints(sa)
	tm := make(map[rune]int)
	for _, v := range t {
		tm[v]++
	}
	ta := getArray(tm)
	sort.Ints(ta)
	if len(sa) != len(ta) {
		fmt.Println("No")
		return
	}
	for i := range sa {
		if sa[i] != ta[i] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
