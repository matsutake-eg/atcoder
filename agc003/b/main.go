package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	iv, _ := strconv.Atoi(sc.Text())
	return iv
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	n := scanInt()
	x := 0
	ans := 0
	for i := 0; i < n; i++ {
		a := scanInt()
		if a > 0 {
			if x > a {
				ans += a
				x = 0
			} else {
				ans += x
				x = a - x
				if x >= 2 {
					ans += x / 2
					x %= 2
				}
			}
		} else {
			x = 0
		}
	}
	fmt.Println(ans)
}
