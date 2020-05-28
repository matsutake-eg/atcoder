package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		s string
		k int
	)
	fmt.Scan(&s, &k)

	wr := bufio.NewWriter(os.Stdout)
	for i := range s {
		d := int(s[i] - 'a')
		if i == len(s)-1 {
			wr.WriteString(string('a' + (d+k)%26))
			break
		}

		if d > 0 && 26-d <= k {
			wr.WriteString("a")
			k -= 26 - d
		} else {
			wr.WriteString(string(s[i]))
		}
	}
	wr.WriteString("\n")
	wr.Flush()
}
