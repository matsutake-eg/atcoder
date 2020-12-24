package main

import (
	"fmt"
	"sort"
)

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	ss := []rune(s)
	sort.Slice(ss, func(i, j int) bool { return ss[i] < ss[j] })
	rt := []rune(t)
	sort.Slice(rt, func(i, j int) bool { return rt[i] > rt[j] })
	if string(ss) < string(rt) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
