package main

import "fmt"

const mod = 2019

func main() {
	var l, r int
	fmt.Scan(&l, &r)

	min := 2020 * 2020
	for i := l; i <= r; i++ {
		for j := i + 1; j <= r; j++ {
			if v := (i % mod) * (j % mod) % mod; v == 0 {
				fmt.Println(0)
				return
			} else if v < min {
				min = v
			}
		}
	}
	fmt.Println(min)
}
