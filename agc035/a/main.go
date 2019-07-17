package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		sc.Scan()
		x, _ := strconv.Atoi(sc.Text())
		m[x]++
	}

	switch len(m) {
	case 1:
		if m[0] == n {
			fmt.Println("Yes")
			return
		}
	case 2:
		if m[0] == n/3 {
			fmt.Println("Yes")
			return
		}
	case 3:
		x := 0
		for k, v := range m {
			if v != n/3 {
				fmt.Println("No")
				return
			}
			x ^= k
		}
		if x == 0 {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
