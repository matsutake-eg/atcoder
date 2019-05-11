package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type City struct {
	Index      int
	Prefecture int
	Year       int
}

type Cities []City

func (cs Cities) Len() int {
	return len(cs)
}

func (cs Cities) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}

func (cs Cities) Less(i, j int) bool {
	if cs[i].Prefecture == cs[j].Prefecture {
		return cs[i].Year < cs[j].Year
	}
	return cs[i].Prefecture < cs[j].Prefecture
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	var cs Cities = make([]City, m)
	for i := range cs {
		cs[i].Index = i
		sc.Scan()
		cs[i].Prefecture, _ = strconv.Atoi(sc.Text())
		sc.Scan()
		cs[i].Year, _ = strconv.Atoi(sc.Text())
	}

	sort.Sort(cs)

	xs := make([]string, m)
	p, c := 0, 0
	for _, v := range cs {
		if p != v.Prefecture {
			p = v.Prefecture
			c = 0
		}
		c++
		xs[v.Index] = fmt.Sprintf("%06d%06d", p, c)
	}

	for _, v := range xs {
		fmt.Println(v)
	}
}
