package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	ts := []string{"dream", "dreamer", "erase", "eraser"}
	var is_found bool

	for {
		is_found = false
		for _, v := range ts {
			if strings.HasSuffix(s, v) {
				is_found = true
				s = strings.TrimSuffix(s, v)
				break
			}
		}

		if !is_found {
			fmt.Println("NO")
			break
		}

		if len(s) == 0 {
			fmt.Println("YES")
			break
		}
	}
}
