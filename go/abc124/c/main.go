package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	sumOdd := 0
	sumEven := 0
	for i, v := range s {
		if i%2 == 0 {
			if v != '0' {
				sumOdd++
			} else {
				sumEven++
			}
		} else {
			if v != '1' {
				sumOdd++
			} else {
				sumEven++
			}
		}
	}
	if sumOdd < sumEven {
		fmt.Println(sumOdd)
	} else {
		fmt.Println(sumEven)
	}
}
