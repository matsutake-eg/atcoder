package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	foundC := false
	for _, r := range s {
		if foundC && r == 'F' {
			fmt.Println("Yes")
			return
		} else if r == 'C' {
			foundC = true
		}
	}
	fmt.Println("No")
}
