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
	a := make([]int, n)
	for i := range a {
		a[i] = scanInt()
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans++
		d := 0
		for j := i + 1; j < n; j++ {
			i = j
			x := a[j] - a[j-1]
			if d == 0 && x != 0 {
				d = x
			}
			if (d > 0 && x < 0) || (d < 0 && x > 0) {
				i = j - 1
				break
			}
		}
	}
	fmt.Println(ans)
}
