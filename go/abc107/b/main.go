package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	a := make([]string, h)
	for i := range a {
		fmt.Scan(&a[i])
	}

	la := make([]bool, h)
	ca := make([]bool, w)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if a[i][j] == '#' {
				la[i] = true
				ca[j] = true
			}
		}
	}

	for i := range a {
		if la[i] {
			for j := range a[i] {
				if ca[j] {
					fmt.Print(string(a[i][j]))
				}
			}
			fmt.Println()
		}
	}
}
