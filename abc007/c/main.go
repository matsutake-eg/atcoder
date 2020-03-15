package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	iv, _ := strconv.Atoi(scanString())
	return iv
}

func scanString() string {
	sc.Scan()
	return sc.Text()
}

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 100001), 100001*100)
}

func main() {
	r, _ := scanInt(), scanInt()
	sy, sx := scanInt(), scanInt()
	gy, gx := scanInt(), scanInt()
	c := make([]string, r)
	for i := 0; i < r; i++ {
		c[i] = scanString()
	}

	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}

	type point struct{ y, x int }
	wk := make(map[point]int)
	start, goal := point{sy - 1, sx - 1}, point{gy - 1, gx - 1}
	ls := list.New()
	ls.PushBack(start)
	wk[start] = 0
	for ls.Len() > 0 {
		p := ls.Remove(ls.Front()).(point)
		if p == goal {
			fmt.Println(wk[p])
			return
		}

		for i := 0; i <= 3; i++ {
			if nx, ny := p.x+dx[i], p.y+dy[i]; nx >= 0 && nx < len(c[0]) && ny >= 0 && ny < r && c[ny][nx] == '.' {
				np := point{ny, nx}
				if _, ok := wk[np]; ok {
					continue
				}
				wk[np] = wk[p] + 1
				ls.PushBack(np)
			}
		}
	}
}
