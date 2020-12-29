package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	x := ""
	ans := 0
	for i := 0; ; {
		if i+1 == len(s) {
			if x != s[i:i+1] {
				ans++
			}
			break
		} else if i == len(s) {
			break
		}

		if x != s[i:i+1] {
			x = s[i : i+1]
			i++
		} else {
			x = s[i : i+2]
			i += 2
		}
		ans++
	}
	fmt.Println(ans)
}
