package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	sumJPY := 0
	sumBTC := 0.0
	var (
		x float64
		u string
	)
	for i := 0; i < n; i++ {
		fmt.Scan(&x, &u)
		if u == "JPY" {
			sumJPY += int(x)
		} else {
			sumBTC += x
		}
	}

	fmt.Println(float64(sumJPY) + sumBTC*380000.0)
}
