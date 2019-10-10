package main

import "fmt"

func main() {
	var (
		s string
		k int
	)
	fmt.Scan(&s, &k)

	cnt := 0
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			cnt++
			i++
		}
	}

	eq := true
	for i := 1; i < len(s); i++ {
		if s[i-1] != s[i] {
			eq = false
			break
		}
	}
	if eq {
		fmt.Println(len(s) * k / 2)
		return
	}

	if s[0] != s[len(s)-1] {
		fmt.Println(k * cnt)
		return
	}

	a := 1
	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			break
		}
		a++
	}
	b := 1
	for i := len(s) - 2; i >= 0; i-- {
		if s[i] != s[len(s)-1] {
			break
		}
		b++
	}
	cnt = 0
	for i := a; i < len(s)-b; i++ {
		if i+1 < len(s)-b && s[i] == s[i+1] {
			cnt++
			i++
		}
	}
	fmt.Println(k*cnt + (k-1)*((a+b)/2) + a/2 + b/2)
}
