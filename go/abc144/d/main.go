package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, x float64
	fmt.Scan(&a, &b, &x)

	var d float64
	if x/(a*a) < b/2 {
		d = a * b * b / (2 * x)
	} else {
		d = 2*b/a - 2*x/(a*a*a)
	}
	fmt.Println(math.Atan(d) / math.Pi * 180)
}
