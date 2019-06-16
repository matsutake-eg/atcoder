package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	d := 0
	cnt := 0
	if d <= x {
		cnt++
	}

	var l int
	for i := 0; i < n; i++ {
		fmt.Scan(&l)
		d = d + l
		if d <= x {
			cnt++
		}
	}
	fmt.Println(cnt)
}
