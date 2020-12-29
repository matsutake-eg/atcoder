package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	n %= 40
	cnt := 0
	for i := 0; i < 2; i++ {
		a, b, d := 1, 21, 1
		if i == 1 {
			a, b, d = 20, 0, -1
		}
		for j := a; j != b; j += d {
			cnt++
			if cnt == n {
				fmt.Println(j)
				return
			}
		}
	}
}
