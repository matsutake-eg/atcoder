package main

import "fmt"

func main() {
	var day string
	fmt.Scan(&day)

	switch day {
	case "Saturday", "Sunday":
		fmt.Println(0)
	case "Monday":
		fmt.Println(5)
	case "Tuesday":
		fmt.Println(4)
	case "Wednesday":
		fmt.Println(3)
	case "Thursday":
		fmt.Println(2)
	case "Friday":
		fmt.Println(1)
	}
}
