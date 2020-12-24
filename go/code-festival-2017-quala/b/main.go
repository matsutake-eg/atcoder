package main

import "fmt"

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)

	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if n*j+m*i-2*i*j == k {
				fmt.Println("Yes")
				return
			}
		}
	}
	fmt.Println("No")
}
