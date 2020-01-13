package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	iv, _ := strconv.Atoi(readStr())
	return iv
}

func readStr() string {
	sc.Scan()
	return sc.Text()
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	h, w := readInt(), readInt()
	s := make([]string, h)
	for i := range s {
		s[i] = readStr()
	}

	type point struct{ x, y int }
	dx := []int{1, 0, -1, 0}
	dy := []int{0, -1, 0, 1}
	ans := 0
	for i, sh := range s {
		for j, v := range sh {
			if v == '#' {
				continue
			}

			qu := list.New()
			qu.PushBack(point{j, i})
			pas := make(map[point]int)
			pas[point{j, i}] = 0
			for {
				p := qu.Remove(qu.Front()).(point)
				for d := 0; d < 4; d++ {
					nx, ny := p.x+dx[d], p.y+dy[d]
					if nx < 0 || nx >= w || ny < 0 || ny >= h || s[ny][nx] == '#' {
						continue
					}
					if _, ok := pas[point{nx, ny}]; ok {
						continue
					}

					qu.PushBack(point{nx, ny})
					mv := pas[p] + 1
					ans = max(ans, mv)
					pas[point{nx, ny}] = mv
				}
				if qu.Len() == 0 {
					break
				}
			}
		}
	}
	fmt.Println(ans)
}
