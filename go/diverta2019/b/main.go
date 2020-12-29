package main

import "fmt"

func main() {
	var r, g, b, n int
	fmt.Scan(&r, &g, &b, &n)

	ans := 0
	for i := 0; i*r <= n; i++ {
		for j := 0; j*g+i*r <= n; j++ {
			if (n-j*g-i*r)%b == 0 {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
