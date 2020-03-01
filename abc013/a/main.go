package main

import "fmt"

func main() {
	var x string
	fmt.Scan(&x)

	switch x {
	case "A":
		fmt.Println(1)
	case "B":
		fmt.Println(2)
	case "C":
		fmt.Println(3)
	case "D":
		fmt.Println(4)
	case "E":
		fmt.Println(5)
	}
}
