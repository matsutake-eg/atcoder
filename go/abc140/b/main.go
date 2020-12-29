package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
		a[i]--
	}
	b := make([]int, n)
	for i := range b {
		fmt.Scan(&b[i])
	}
	c := make([]int, n-1)
	for i := range c {
		fmt.Scan(&c[i])
	}

	old := -2
	ans := 0
	for _, v := range a {
		ans += b[v]
		if v-old == 1 {
			ans += c[old]
		}
		old = v
	}
	fmt.Println(ans)
}
