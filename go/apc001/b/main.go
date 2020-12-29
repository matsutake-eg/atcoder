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
	b := make([]int, n)
	for i := range b {
		b[i] = scanInt()
	}

	d := 0
	cnt := 0
	for i := 0; i < n; i++ {
		d += b[i] - a[i]
		if a[i] < b[i] {
			cnt += (b[i] - a[i] + 1) / 2
		}
	}
	if d >= cnt {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
