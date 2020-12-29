package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	a, r := x/11, x%11
	ans := a * 2
	if r >= 1 && r <= 6 {
		ans++
	} else if r >= 7 {
		ans += 2
	}
	fmt.Println(ans)
}
