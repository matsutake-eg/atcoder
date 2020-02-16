package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	n := make([]int, 4)
	for i := range n {
		n[i] = int(s[i] - '0')
	}
	op := []string{"+++", "-++", "+-+", "--+", "++-", "-+-", "+--", "---"}
	for _, o := range op {
		ans := n[0]
		for i, r := range o {
			if r == '+' {
				ans += n[i+1]
			} else {
				ans += -n[i+1]
			}
		}
		if ans == 7 {
			fmt.Printf("%d%s%d%s%d%s%d=7\n", n[0], o[0:1], n[1], o[1:2], n[2], o[2:3], n[3])
			return
		}
	}
}
