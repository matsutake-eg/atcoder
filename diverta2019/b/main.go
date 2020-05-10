package main

import "fmt"

func main() {
	var r, g, b, n int
	fmt.Scan(&r, &g, &b, &n)

	ans := 0
	for i := 0; i*r <= n; i++ {
		for j := 0; i*r+j*g <= n; j++ {
			if (n-i*r-j*g)%b == 0 {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
