package main

import (
	"fmt"
	"math"
)

func main() {
	var x int
	fmt.Scan(&x)

	m := 100
	for i := 0; ; i++ {
		if m >= x {
			fmt.Println(i)
			return
		}
		m = int(math.Floor(float64(m) * 1.01))
	}
}
