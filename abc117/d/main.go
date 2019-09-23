package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	bc := 0
	for i := 1; i <= k; i <<= 1 {
		bc++
	}

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	a := make([]int, n)
	bits := make([]int, bc)
	for i := range a {
		sc.Scan()
		a[i], _ = strconv.Atoi(sc.Text())
		for j := range bits {
			if a[i]&(1<<uint(j)) == 0 {
				bits[j]++
			}
		}
	}

	x := 0
	for i := len(bits) - 1; i >= 0; i-- {
		if v := 1 << uint(i); bits[i] > len(a)/2 && x+v <= k {
			x += v
		}
	}

	ans := 0
	for _, v := range a {
		ans += x ^ v
	}
	fmt.Println(ans)
}
