package main

import "fmt"

func main() {
	var sa, sb, sc string
	fmt.Scan(&sa, &sb, &sc)

	s := sa[0:1]
	for {
		switch s {
		case "a":
			if len(sa) == 0 {
				fmt.Println("A")
				return
			}
			s = sa[0:1]
			sa = sa[1:]
		case "b":
			if len(sb) == 0 {
				fmt.Println("B")
				return
			}
			s = sb[0:1]
			sb = sb[1:]
		case "c":
			if len(sc) == 0 {
				fmt.Println("C")
				return
			}
			s = sc[0:1]
			sc = sc[1:]
		}
	}
}
