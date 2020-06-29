package main

import (
	"fmt"
	"math"
)

func main() {
	var h, b float64
	fmt.Scan(&h, &b)

	fmt.Println(b * math.Pow((h/1000.0), 2) * 100)
}
