package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	a := make([]int, 4)
	for i, r := range s {
		a[i] = int(r - '0')
	}

	switch {
	case a[0]+a[1]+a[2]+a[3] == 7:
		fmt.Printf("%v+%v+%v+%v=7\n", a[0], a[1], a[2], a[3])
	case a[0]+a[1]+a[2]-a[3] == 7:
		fmt.Printf("%v+%v+%v-%v=7\n", a[0], a[1], a[2], a[3])
	case a[0]+a[1]-a[2]+a[3] == 7:
		fmt.Printf("%v+%v-%v+%v=7\n", a[0], a[1], a[2], a[3])
	case a[0]+a[1]-a[2]-a[3] == 7:
		fmt.Printf("%v+%v-%v-%v=7\n", a[0], a[1], a[2], a[3])
	case a[0]-a[1]+a[2]+a[3] == 7:
		fmt.Printf("%v-%v+%v+%v=7\n", a[0], a[1], a[2], a[3])
	case a[0]-a[1]-a[2]+a[3] == 7:
		fmt.Printf("%v-%v-%v+%v=7\n", a[0], a[1], a[2], a[3])
	case a[0]-a[1]+a[2]-a[3] == 7:
		fmt.Printf("%v-%v+%v-%v=7\n", a[0], a[1], a[2], a[3])
	case a[0]-a[1]-a[2]-a[3] == 7:
		fmt.Printf("%v-%v-%v-%v=7\n", a[0], a[1], a[2], a[3])
	}
}
