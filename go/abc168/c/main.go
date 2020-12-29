package main

import (
	"fmt"
	"math"
)

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

func main() {
	var a, b, h, m float64
	fmt.Scan(&a, &b, &h, &m)

	hang := 30*h + 30*(m/60)
	mang := 360 * (m / 60)
	ang := abs(hang - mang)
	ang = min(ang, 360-ang)
	fmt.Println(math.Sqrt(a*a + b*b - 2*a*b*math.Cos(math.Pi*ang/180)))
}
