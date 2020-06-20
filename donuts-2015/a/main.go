package main

import (
	"fmt"
	"math"
)

func main() {
	var r, d float64
	fmt.Scan(&r, &d)

	fmt.Println(r * r * math.Pi * d * 2 * math.Pi)
}
