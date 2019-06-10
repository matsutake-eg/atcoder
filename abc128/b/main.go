package main

import (
	"fmt"
	"sort"
)

type Restaurant struct {
	Name  string
	Point int
	Index int
}
type Restaurants []Restaurant

func (r Restaurants) Len() int      { return len(r) }
func (r Restaurants) Swap(i, j int) { r[i], r[j] = r[j], r[i] }
func (r Restaurants) Less(i, j int) bool {
	if r[i].Name == r[j].Name {
		return r[i].Point > r[j].Point
	}
	return r[i].Name < r[j].Name
}

func main() {
	var n int
	fmt.Scan(&n)
	var rs Restaurants = make([]Restaurant, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&rs[i].Name, &rs[i].Point)
		rs[i].Index = i + 1
	}

	sort.Sort(rs)
	for _, v := range rs {
		fmt.Println(v.Index)
	}
}
