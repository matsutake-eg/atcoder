package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	s := make([]string, h)
	for i := range s {
		fmt.Scan(&s[i])
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == '.' {
				continue
			}
			bad := true
			if i > 0 && s[i-1][j] == '#' {
				bad = false
			}
			if i < h-1 && s[i+1][j] == '#' {
				bad = false
			}
			if j > 0 && s[i][j-1] == '#' {
				bad = false
			}
			if j < w-1 && s[i][j+1] == '#' {
				bad = false
			}
			if bad {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
}
