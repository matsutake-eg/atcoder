package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	ans := 0
	switch x {
	case 1:
		ans += 300000
	case 2:
		ans += 200000
	case 3:
		ans += 100000
	}
	switch y {
	case 1:
		ans += 300000
	case 2:
		ans += 200000
	case 3:
		ans += 100000
	}
	if x == 1 && y == 1 {
		ans += 400000
	}
	fmt.Println(ans)
}
