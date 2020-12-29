package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)

	switch n[len(n)-1] {
	case '2', '4', '5', '7', '9':
		fmt.Println("hon")
	case '0', '1', '6', '8':
		fmt.Println("pon")
	case '3':
		fmt.Println("bon")
	}
}
