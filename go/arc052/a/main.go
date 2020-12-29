package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	wr := bufio.NewWriter(os.Stdout)
	for _, r := range s {
		sr := string(r)
		for i := 0; i <= 9; i++ {
			if strings.Contains(sr, strconv.Itoa(i)) {
				wr.WriteString(sr)
				break
			}
		}
	}
	wr.WriteString("\n")
	wr.Flush()
}
