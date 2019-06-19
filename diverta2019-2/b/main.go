package main

import (
	"fmt"
)

type Point struct{ X, Y int }
type Key struct{ Dx, Dy int }

func main() {
	var n int
	fmt.Scan(&n)

	if n == 1 {
		fmt.Println(1)
		return
	}

	ps := make([]Point, n)
	for i := range ps {
		fmt.Scan(&ps[i].X, &ps[i].Y)
	}

	xm := make(map[Key]int)
	for _, v1 := range ps {
		for _, v2 := range ps {
			if v1 == v2 {
				continue
			}
			xm[Key{v1.X - v2.X, v1.Y - v2.Y}]++
		}
	}

	max := 0
	for _, v := range xm {
		if v > max {
			max = v
		}
	}
	fmt.Println(n - max)
}
