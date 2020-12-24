package main

import "fmt"

func main() {
	var m, d int
	fmt.Scan(&m, &d)

	ans := 0
	for i := 11; i <= d; i++ {
		d1, d2 := i%10, i/10
		if d1 == 1 || d2 == 1 {
			continue
		}
		for j := 2; j <= m; j++ {
			if j == d1*d2 {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
