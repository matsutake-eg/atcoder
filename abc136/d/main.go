package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)

	ans := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == 'R' {
			j := 0
			for s[i+j+1] == 'R' {
				j++
			}
			ans[i+j] += (j + 1) - (j+1)/2
			ans[i+j+1] += (j + 1) / 2
			i += j
		} else {
			j := 0
			for i+j+1 < len(s) && s[i+j+1] == 'L' {
				j++
			}
			ans[i] += (j + 1) - (j+1)/2
			ans[i-1] += (j + 1) / 2
			i += j
		}
	}

	wr := bufio.NewWriter(os.Stdout)
	for i, v := range ans {
		wr.WriteString(strconv.Itoa(v))
		if i < len(ans)-1 {
			wr.WriteString(" ")
		} else {
			wr.WriteString("\n")
		}
	}
	wr.Flush()
}
