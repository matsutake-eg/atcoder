package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	for {
		if a+1 < b {
			a++
		} else if a-1 >= 1 {
			a--
		}
		if b-1 > a {
			b--
		} else if b+1 <= n {
			b++
		}

		if a == 1 {
			fmt.Println("Borys")
			return
		}
		if b == n {
			fmt.Println("Alice")
			return
		}
	}
}
