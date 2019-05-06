package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type City struct {
	Index        int
	PrefectureNo int
	Year         int
	ServiceNo    string
}

type By func(c1, c2 *City) bool

func (by By) Sort(cities []City) {
	cs := &citySorter{
		cities: cities,
		by:     by,
	}
	sort.Sort(cs)
}

type citySorter struct {
	cities []City
	by     func(c1, c2 *City) bool
}

func (c *citySorter) Len() int {
	return len(c.cities)
}

func (c *citySorter) Swap(i, j int) {
	c.cities[i], c.cities[j] = c.cities[j], c.cities[i]
}

func (c *citySorter) Less(i, j int) bool {
	return c.by(&c.cities[i], &c.cities[j])
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	cs := make([]City, m)
	for i := range cs {
		cs[i].Index = i
		sc.Scan()
		cs[i].PrefectureNo, _ = strconv.Atoi(sc.Text())
		sc.Scan()
		cs[i].Year, _ = strconv.Atoi(sc.Text())
	}

	By(func(c1, c2 *City) bool {
		if c1.PrefectureNo == c2.PrefectureNo {
			return c1.Year < c2.Year
		}
		return c1.PrefectureNo < c2.PrefectureNo
	}).Sort(cs)
	p := 0
	var y int
	for i, v := range cs {
		if p != v.PrefectureNo {
			p = v.PrefectureNo
			y = 0
		}
		y++
		cs[i].ServiceNo = fmt.Sprintf("%06d%06d", p, y)
	}

	By(func(c1, c2 *City) bool {
		return c1.Index < c2.Index
	}).Sort(cs)
	for _, v := range cs {
		fmt.Println(v.ServiceNo)
	}
}
