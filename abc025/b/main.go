package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	iv, _ := strconv.Atoi(scanString())
	return iv
}

func scanString() string {
	sc.Scan()
	return sc.Text()
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	n, a, b := scanInt(), scanInt(), scanInt()
	pos := 0
	for i := 0; i < n; i++ {
		s, d := scanString(), scanInt()
		if s == "East" {
			if d < a {
				pos += a
			} else if d > b {
				pos += b
			} else {
				pos += d
			}
		} else {
			if d < a {
				pos -= a
			} else if d > b {
				pos -= b
			} else {
				pos -= d
			}
		}
	}

	if pos == 0 {
		fmt.Println(0)
	} else if pos > 0 {
		fmt.Println("East", pos)
	} else {
		fmt.Println("West", -pos)
	}
}
