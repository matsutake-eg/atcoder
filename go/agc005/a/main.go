package main

import "fmt"

func main() {
	var x string
	fmt.Scan(&x)

	ans := len(x)
	cnt := 0
	for _, r := range x {
		if r == 'S' {
			cnt++
		} else if r == 'T' && cnt > 0 {
			cnt--
			ans -= 2
		}
	}
	fmt.Println(ans)
}
