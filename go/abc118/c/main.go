package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var A []int

func solve(x int) int {
	for _, v := range A {
		if r := v % x; r != 0 {
			return solve(r)
		}
	}
	return x
}

func main() {
	var n int
	fmt.Scan(&n)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	A = make([]int, n)
	min := 1000000001
	for i := range A {
		sc.Scan()
		A[i], _ = strconv.Atoi(sc.Text())
		if A[i] < min {
			min = A[i]
		}
	}

	fmt.Println(solve(min))
}
