package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	p := make([]int, n)
	for i := range p {
		fmt.Scan(&p[i])
	}

	ans := 0
	for i := 1; i < n-1; i++ {
		switch {
		case p[i-1] == p[i] && p[i] == p[i+1]:
			continue
		case p[i] > p[i-1] && p[i] > p[i+1]:
			continue
		case p[i] < p[i-1] && p[i] < p[i+1]:
			continue
		}
		ans++
	}
	fmt.Println(ans)
}
