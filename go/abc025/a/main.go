package main

import "fmt"

func main() {
	var (
		s string
		n int
	)
	fmt.Scan(&s, &n)

	cnt := 0
	for _, r1 := range s {
		for _, r2 := range s {
			cnt++
			if cnt == n {
				fmt.Println(string(r1) + string(r2))
				return
			}
		}
	}
}
