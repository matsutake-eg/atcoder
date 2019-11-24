package main

import (
	"container/list"
	"fmt"
)

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	s := make([]string, h)
	ans := 0
	passed := make([][]bool, h)
	for i := range passed {
		passed[i] = make([]bool, w)
	}
	for i := range s {
		fmt.Scan(&s[i])
		for j, r := range s[i] {
			if r == '.' {
				ans++
				continue
			}
			passed[i][j] = true
		}
	}

	type point struct{ x, y, d int }
	ps := list.New()
	ps.PushBack(point{0, 0, 1})

	dx := [4]int{1, 0, -1, 0}
	dy := [4]int{0, 1, 0, -1}
	for {
		p := ps.Remove(ps.Front()).(point)
		if p.x == h-1 && p.y == w-1 {
			ans -= p.d
			break
		}

		for i := 0; i < 4; i++ {
			if nx, ny := p.x+dx[i], p.y+dy[i]; nx >= 0 && nx < h && ny >= 0 && ny < w && !passed[nx][ny] {
				ps.PushBack(point{nx, ny, p.d + 1})
				passed[nx][ny] = true
			}
		}

		if ps.Len() == 0 {
			fmt.Println(-1)
			return
		}
	}
	fmt.Println(ans)
}
