package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	ans := ""
	for n != 0 {
		if n%-2 == 0 {
			ans = "0" + ans
		} else {
			ans = "1" + ans
			n--
		}
		n /= -2
	}
	if ans == "" {
		ans = "0"
	}
	fmt.Println(ans)
}
