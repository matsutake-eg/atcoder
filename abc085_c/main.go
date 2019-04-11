package main

import "fmt"

func main() {
	var n, y int
	fmt.Scan(&n, &y)

	is_answer_exist := false
	for i := 0; i <= n; i++ {
		if is_answer_exist {
			break
		}
		for j := 0; i+j <= n; j++ {
			if is_answer_exist {
				break
			}
			if i*10000+j*5000+(n-i-j)*1000 == y {
				is_answer_exist = true
				fmt.Println(i, j, n-i-j)
			}
		}
	}

	if !is_answer_exist {
		fmt.Println("-1 -1 -1")
	}
}
