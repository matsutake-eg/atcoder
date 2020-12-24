package main

import "fmt"

func main() {
	var h int
	fmt.Scan(&h)

	cnt := 1
	ans := 0
	for h > 0 {
		h /= 2
		ans += cnt
		cnt *= 2
	}
	fmt.Println(ans)
}
