package main

import "fmt"

func main() {
	var d, n int
	fmt.Scan(&d, &n)

	b := 1
	for i := 0; i < d; i++ {
		b *= 100
	}
	if n == 100 {
		n = 101
	}
	fmt.Println(b * n)
}
