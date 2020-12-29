package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int
	fmt.Scan(&a)

	for i := 10; ; i++ {
		s := strconv.Itoa(i)
		sum := 0
		d := 1
		for j := len(s) - 1; j >= 0; j-- {
			sum += int(s[j]-'0') * d
			d *= i
		}
		if sum == a {
			fmt.Println(i)
			return
		} else if sum/a >= 1 {
			break
		}
	}
	fmt.Println(-1)
}
