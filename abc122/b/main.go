package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	count := 0
	max := 0
	for _, r := range s {
		if r == 'A' || r == 'C' || r == 'G' || r == 'T' {
			count++
		} else {
			if count > max {
				max = count
			}
			count = 0
		}
	}
	if count > max {
		max = count
	}
	fmt.Println(max)
}
