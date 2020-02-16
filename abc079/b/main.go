package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	l := make([]int, n+1)
	l[0], l[1] = 2, 1
	for i := 2; i <= n; i++ {
		l[i] = l[i-1] + l[i-2]
	}
	fmt.Println(l[n])
}
