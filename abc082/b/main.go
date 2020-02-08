package main

import (
	"fmt"
	"sort"
)

type sortString []rune

func (s sortString) Len() int           { return len(s) }
func (s sortString) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortString) Less(i, j int) bool { return s[i] < s[j] }

type sortStringR []rune

func (s sortStringR) Len() int           { return len(s) }
func (s sortStringR) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortStringR) Less(i, j int) bool { return s[i] > s[j] }

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	ss := sortString(s)
	sort.Sort(ss)
	rt := sortStringR(t)
	sort.Sort(rt)
	if string(ss) < string(rt) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
