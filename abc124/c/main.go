package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	var sumOdd, sumEven int
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
