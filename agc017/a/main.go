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

func cmb(x, y int) int {
	if y > x {
		return 0
	}

	nu, de := 1, 1
	for i := 0; i < y; i++ {
		nu = nu * (x - i)
		de = de * (i + 1)
		if nu%de == 0 {
			nu /= de
			de = 1
		}
	}
	return nu / de
}

func main() {
	n, p := scanInt(), scanInt()
	e, o := 0, 0
	for i := 0; i < n; i++ {
		a := scanInt()
		if a%2 == 0 {
			e++
		} else {
			o++
		}
	}
	ans := 0
	for i := 0; i <= n; i++ {
		for j := p; j <= i; j += 2 {
			ans += cmb(o, j) * cmb(e, i-j)
		}
	}
	fmt.Println(ans)
}
