package main

import "fmt"

func main() {
	var b string
	fmt.Scan(&b)
	switch b {
	case "A":
		b = "T"
	case "T":
		b = "A"
	case "C":
		b = "G"
	case "G":
		b = "C"
	}
	fmt.Println(b)
}
