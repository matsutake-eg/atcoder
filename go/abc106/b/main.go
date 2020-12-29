package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	ans := 0
	for i := 1; i <= n; i++ {
		cnt := 0
		for j := 1; j <= i; j++ {
			if j%2 == 1 && i%j == 0 {
				cnt++
			}
		}
		if cnt == 8 {
			ans++
		}
	}
	fmt.Println(ans)
}
