package main

import "fmt"

func main() {
	var (
		n       int
		a, b, c string
	)
	fmt.Scan(&n, &a, &b, &c)

	ans := 0
	for i := 0; i < n; i++ {
		if a[i] != b[i] && b[i] != c[i] && c[i] != a[i] {
			ans += 2
		} else if !(a[i] == b[i] && b[i] == c[i] && c[i] == a[i]) {
			ans++
		}
	}
	fmt.Println(ans)
}
