package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)

	cnt := 0
	gd := false
	for i := range n[:len(n)-1] {
		if n[i] == n[i+1] {
			cnt++
			if cnt >= 2 {
				gd = true
			}
		} else {
			cnt = 0
		}
	}
	if gd {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
