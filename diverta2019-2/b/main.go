package main

import (
	"fmt"
	"strconv"
)

type Point struct{ X, Y int }

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

	xm := make(map[string]int)
	for _, v1 := range ps {
		for _, v2 := range ps {
			dx := strconv.Itoa(v1.X - v2.X)
			dy := strconv.Itoa(v1.Y - v2.Y)
			xm[dx+"_"+dy]++
		}
	}
	xm["0_0"] = 1

	max := 0
	for _, v := range xm {
		if v > max {
			max = v
		}
	}
	fmt.Println(n - max)
}
