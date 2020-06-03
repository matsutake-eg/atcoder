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
	b := make([]int, n)
	for i := range a {
		a[i] = scanInt()
	}

	for i := range a {
		if i > 0 && a[i-1] == a[i] {
			b[i]++
		}
		if i < n-1 && a[i] == a[i+1] {
			b[i]++
		}
	}

	ans := 0
	for i := 1; i < n-1; i++ {
		if b[i] == 2 {
			ans++
			b[i] = 0
			b[i-1]--
			b[i+1]--
		}
	}
	for i := 0; i < n-1; i++ {
		if b[i] == 1 {
			ans++
			b[i] = 0
			b[i+1] = 0
		}
	}
	fmt.Println(ans)
}
