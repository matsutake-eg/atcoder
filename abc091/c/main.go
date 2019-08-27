package main

import (
	"fmt"
	"sort"
)

type point struct{ x, y int }
type sortByX []point

func (s sortByX) Len() int           { return len(s) }
func (s sortByX) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortByX) Less(i, j int) bool { return s[i].x < s[j].x }

type sortByY []point

func (s sortByY) Len() int           { return len(s) }
func (s sortByY) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortByY) Less(i, j int) bool { return s[i].y > s[j].y }

func main() {
	var n int
	fmt.Scan(&n)

	rs := make(map[point]bool, n)
	for i := 0; i < n; i++ {
		p := point{}
		fmt.Scan(&p.x, &p.y)
		rs[p] = true
	}
	bs := sortByX(make([]point, n))
	for i := range bs {
		fmt.Scan(&bs[i].x, &bs[i].y)
	}

	sort.Sort(bs)
	ans := 0
	for _, b := range bs {
		ps := sortByY(make([]point, 0, n))
		for r := range rs {
			if r.x < b.x && r.y < b.y {
				ps = append(ps, r)
			}
		}

		if len(ps) > 0 {
			sort.Sort(ps)
			delete(rs, ps[0])
			ans++
		}
	}
	fmt.Println(ans)
}
