package main

import "fmt"

func main() {
	var x string
	fmt.Scan(&x)

	if x == "" {
		fmt.Println("YES")
		return
	}

	for i := 0; i < len(x); i++ {
		switch x[i] {
		case 'o':
			continue
		case 'k':
			continue
		case 'u':
			continue
		case 'c':
			if i+1 < len(x) && x[i+1] == 'h' {
				i++
				continue
			}
		}
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
}
