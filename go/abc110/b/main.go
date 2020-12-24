package main

import "fmt"

func main() {
	var N, M, X, Y int
	fmt.Scan(&N, &M, &X, &Y)
	var (
		l = X
		t int
	)
	for i := 0; i < N; i++ {
		fmt.Scan(&t)
		if t > l {
			l = t
		}
	}
	r := Y
	for i := 0; i < M; i++ {
		fmt.Scan(&t)
		if t < r {
			r = t
		}
	}
	if r-l >= 1 {
		fmt.Println("No War")
	} else {
		fmt.Println("War")
	}
}
