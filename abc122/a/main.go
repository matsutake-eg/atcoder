package main

import "fmt"

func main() {
	var b string
	fmt.Scan(&b)

	switch b {
	case "A":
		fmt.Println("T")
	case "C":
		fmt.Println("G")
	case "G":
		fmt.Println("C")
	case "T":
		fmt.Println("A")
	}
}
