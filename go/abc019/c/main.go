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
	xm := make(map[int]bool)
	for i := 0; i < n; i++ {
		a := scanInt()
		xm[a] = true
	}

	ans := 0
	for a := range xm {
		if a%2 == 1 {
			ans++
			continue
		}

		found := false
		for x := a; ; {
			x /= 2
			if x%2 == 1 {
				if xm[x] {
					found = true
				}
				break
			} else {
				if xm[x] {
					found = true
					break
				}
			}
		}
		if !found {
			ans++
		}
	}
	fmt.Println(ans)
}
