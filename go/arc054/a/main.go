package main

import (
	"fmt"
)

func min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

func main() {
	var l, x, y, s, d int
	fmt.Scan(&l, &x, &y, &s, &d)

	if d >= s {
		k := d - s
		if x >= y {
			fmt.Println(float64(k) / float64(x+y))
		} else {
			fmt.Println(min(float64(k)/float64(x+y), float64(l-k)/float64(y-x)))
		}
	} else {
		k := s - d
		if x >= y {
			fmt.Println(float64(l-k) / float64(x+y))
		} else {
			fmt.Println(min(float64(l-k)/float64(x+y), float64(k)/float64(y-x)))
		}
	}
}
