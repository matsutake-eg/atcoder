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
	sc.Buffer(make([]byte, 100001), 100001*100)
}

func main() {
	h, w := scanInt(), scanInt()
	a := make([]string, h)
	for i := range a {
		a[i] = scanString()
	}

	for i := range a {
		for j := range a[i] {
			if a[i][j] == '.' {
				continue
			}

			from := 0
			if i-1 >= 0 && a[i-1][j] == '#' {
				from++
			}
			if j-1 >= 0 && a[i][j-1] == '#' {
				from++
			}
			if i == 0 && j == 0 {
				from = 1
			}
			to := 0
			if i+1 < h && a[i+1][j] == '#' {
				to++
			}
			if j+1 < w && a[i][j+1] == '#' {
				to++
			}
			if i == h-1 && j == w-1 {
				to = 1
			}
			if !(from == 1 && to == 1) {
				fmt.Println("Impossible")
				return
			}
		}
	}
	fmt.Println("Possible")
}
