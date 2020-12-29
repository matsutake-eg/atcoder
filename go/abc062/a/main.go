package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	if x == 2 || y == 2 {
		fmt.Println("No")
		return
	}

	one := map[int]bool{1: true, 3: true, 5: true, 7: true, 8: true, 10: true, 12: true}
	two := map[int]bool{4: true, 6: true, 9: true, 11: true}
	if (one[x] && one[y]) || (two[x] && two[y]) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
