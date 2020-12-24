package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	p := make([]int, n)
	for i := range p {
		fmt.Scan(&p[i])
	}

	cnt := 0
	for i := range p {
		v := p[i]
		if v == i+1 {
			continue
		}
		cnt++
	}
	if cnt == 0 || cnt == 2 {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}
