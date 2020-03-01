package main

import "fmt"

func main() {
	var (
		n int
		s string
	)
	fmt.Scan(&n, &s)

	ac := "b"
	i := 0
	for {
		if ac == s {
			fmt.Println(i)
			return
		} else if len(ac) > n {
			fmt.Println(-1)
			return
		}

		i++
		switch i % 3 {
		case 1:
			ac = "a" + ac + "c"
		case 2:
			ac = "c" + ac + "a"
		case 0:
			ac = "b" + ac + "b"
		}
	}
}
