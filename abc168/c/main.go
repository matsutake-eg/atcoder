package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, h, m float64
	fmt.Scan(&a, &b, &h, &m)

	if h >= 12 {
		h -= 12
	}

	l := 360 * (float64(m) / 60)
	s := 30*float64(h) + 30*(float64(m)/60)
	var ang float64
	if l > s {
		ang = l - s
	} else {
		ang = s - l
	}
	if ang >= 180 {
		ang = 360 - ang
	}
	fmt.Println(math.Sqrt(a*a + b*b - 2*a*b*math.Cos(math.Pi*ang/180)))
}
