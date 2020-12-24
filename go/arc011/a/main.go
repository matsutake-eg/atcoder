package main

import "fmt"

func main() {
	var m, n, N int
	fmt.Scan(&m, &n, &N)

	ans := N
	r := 0
	for N > 0 {
		N += r
		r = N % m
		N = N / m * n
		ans += N
	}
	fmt.Println(ans)
}
