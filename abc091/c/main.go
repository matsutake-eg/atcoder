package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	iv, _ := strconv.Atoi(sc.Text())
	return iv
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	var n int
	fmt.Scan(&n)

	type point struct{ x, y int }
	rs := make(map[point]bool)
	for i := 0; i < n; i++ {
		x, y := scanInt(), scanInt()
		rs[point{x, y}] = true
	}

	bs := make([]point, n)
	for i := range bs {
		bs[i].x, bs[i].y = scanInt(), scanInt()
	}

	sort.Slice(bs, func(i, j int) bool { return bs[i].x < bs[j].x })
	ans := 0
	for _, b := range bs {
		ps := make([]point, 0, n)
		for r := range rs {
			if r.x < b.x && r.y < b.y {
				ps = append(ps, r)
			}
		}

		if len(ps) > 0 {
			sort.Slice(ps, func(i, j int) bool { return ps[i].y > ps[j].y })
			delete(rs, ps[0])
			ans++
		}
	}
	fmt.Println(ans)
}
