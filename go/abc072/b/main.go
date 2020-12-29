package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	ans := make([]byte, 0, len(s)/2+1)
	for i := 0; i < len(s); i += 2 {
		ans = append(ans, s[i])
	}
	fmt.Println(string(ans))
}
