package main

import "fmt"

func main() {
	var m float64
	fmt.Scan(&m)

	m /= 1000
	if m < 0.1 {
		fmt.Println("00")
	} else if m >= 0.1 && m <= 5 {
		fmt.Printf("%02v\n", int(m*10))
	} else if m >= 6 && m <= 30 {
		fmt.Println(m + 50)
	} else if m >= 35 && m <= 70 {
		fmt.Println((m-30)/5 + 80)
	} else if m > 70 {
		fmt.Println(89)
	}
}
