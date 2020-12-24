package main

import "fmt"

func main() {
	var (
		n, l int
		s    string
	)
	fmt.Scan(&n, &l, &s)

	ans := 0
	cnt := 0
	for _, r := range s {
		switch r {
		case '+':
			cnt++
		case '-':
			if cnt > 0 {
				cnt--
			}
		}
		if cnt == l {
			ans++
			cnt = 0
		}
	}
	fmt.Println(ans)
}
