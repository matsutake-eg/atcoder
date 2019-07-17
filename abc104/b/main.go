package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	if s[0] != 'A' {
		fmt.Println("WA")
		return
	}

	cnt := 0
	for _, r := range s[2 : len(s)-1] {
		if r == 'C' {
			cnt++
		}
	}
	if cnt != 1 {
		fmt.Println("WA")
		return

	}

	cntA, cntC, cntCL := 0, 0, 0
	for _, r := range s {
		switch {
		case r == 'A':
			cntA++
		case r == 'C':
			cntC++
		case r <= 90:
			cntCL++
		}
	}
	if cntA == 1 && cntC == 1 && cntCL == 0 {
		fmt.Println("AC")
		return
	}
	fmt.Println("WA")
}
