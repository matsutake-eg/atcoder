package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	for i, r := range s {
		if i%2 == 0 {
			switch r {
			case 'R', 'U', 'D':
			default:
				fmt.Println("No")
				return
			}
		} else {
			switch r {
			case 'L', 'U', 'D':
			default:
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
}
