package main

import (
	"fmt"
	"sort"
)

type Restaurant struct {
	name  string
	point int
	inedx int
}
type Restaurants []Restaurant

func (r Restaurants) Len() int      { return len(r) }
func (r Restaurants) Swap(i, j int) { r[i], r[j] = r[j], r[i] }
func (r Restaurants) Less(i, j int) bool {
	if r[i].name == r[j].name {
		return r[i].point > r[j].point
	}
	return r[i].name < r[j].name
}

func main() {
	var n int
	fmt.Scan(&n)
	var rs Restaurants = make([]Restaurant, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&rs[i].name, &rs[i].point)
		rs[i].inedx = i + 1
	}
	sort.Sort(rs)
	for _, v := range rs {
		fmt.Println(v.inedx)
	}
}
