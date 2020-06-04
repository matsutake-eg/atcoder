package main

import "fmt"

func main() {
	var (
		n, a, b, c, d int
		s             string
	)
	fmt.Scan(&n, &a, &b, &c, &d, &s)

	for i := b - 1; i <= d-1; i++ {
		if s[i] == '#' && i+1 <= d-1 && s[i+1] == '#' {
			fmt.Println("No")
			return
		}
	}
	for i := a - 1; i <= c-1; i++ {
		if s[i] == '#' && i+1 <= c-1 && s[i+1] == '#' {
			fmt.Println("No")
			return
		}
	}

	if c < d {
		fmt.Println("Yes")
	} else {
		for i := b - 1; i <= d-1; i++ {
			if s[i-1:i+2] == "..." {
				fmt.Println("Yes")
				return
			}
		}
		fmt.Println("No")
	}
}
