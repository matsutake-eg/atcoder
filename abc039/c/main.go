package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	switch s[:12] {
	case "WBWBWWBWBWBW":
		fmt.Println("Do")
	case "WBWWBWBWBWWB":
		fmt.Println("Re")
	case "WWBWBWBWWBWB":
		fmt.Println("Mi")
	case "WBWBWBWWBWBW":
		fmt.Println("Fa")
	case "WBWBWWBWBWWB":
		fmt.Println("So")
	case "WBWWBWBWWBWB":
		fmt.Println("La")
	case "WWBWBWWBWBWB":
		fmt.Println("Si")
	}
}
