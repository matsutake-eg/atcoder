package main

import "fmt"

func calc(y, m, d int) int {
	return 365*y + y/4 - y/100 + y/400 + (306 * (m + 1) / 10) + d - 429
}

func main() {
	var y, m, d int
	fmt.Scan(&y, &m, &d)

	if m == 1 || m == 2 {
		y--
		m += 12
	}
	fmt.Println(calc(2014, 5, 17) - calc(y, m, d))
}
