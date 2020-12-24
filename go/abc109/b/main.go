package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	wa := make([]string, n)
	wm := make(map[string]int)
	for i := range wa {
		fmt.Scan(&wa[i])
		wm[wa[i]] = 0
	}
	if len(wa) != len(wm) {
		fmt.Println("No")
		return
	}
	for i := 0; i < n-1; i++ {
		w1, w2 := wa[i], wa[i+1]
		if w1[len(w1)-1] != w2[0] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
