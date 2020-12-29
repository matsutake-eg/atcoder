package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	words := []string{"dream", "dreamer", "erase", "eraser"}
	for i := len(s); i > 0; {
		found := false
		for _, w := range words {
			if ni := i - len(w); ni >= 0 && w == s[ni:i] {
				found = true
				i = ni
				break
			}
		}
		if !found {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
