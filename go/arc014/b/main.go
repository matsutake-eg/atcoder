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
	sc.Buffer(make([]byte, 100000), 100000000)
}

func main() {
	n := scanInt()
	bw := scanString()
	xm := make(map[string]bool)
	xm[bw] = true
	for i := 1; i < n; i++ {
		w := scanString()
		if i%2 == 1 {
			if xm[w] {
				fmt.Println("WIN")
				return
			}
		} else {
			if xm[w] {
				fmt.Println("LOSE")
				return
			}
		}
		xm[w] = true

		if bw[len(bw)-1] != w[0] {
			if i%2 == 1 {
				fmt.Println("WIN")
			} else {
				fmt.Println("LOSE")
			}
			return
		}
		bw = w
	}
	fmt.Println("DRAW")
}
