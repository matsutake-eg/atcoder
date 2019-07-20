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
	a := make([]int, n)
	max1, max2 := 0, 0
	for i := 0; i < n; i++ {
		sc.Scan()
		a[i], _ = strconv.Atoi(sc.Text())
		if a[i] > max1 {
			max1, max2 = a[i], max1
			continue
		}
		if a[i] > max2 {
			max2 = a[i]
		}
	}

	for _, v := range a {
		if v == max1 {
			fmt.Println(max2)
			continue
		}
		fmt.Println(max1)
	}
}
