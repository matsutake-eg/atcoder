package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	ans := ""
	for n > 0 {
		n--
		ans = string('a'+n%26) + ans
		n /= 26
	}
	fmt.Println(ans)
}
