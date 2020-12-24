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

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	ans := 0
	for i := 0; i < n; i++ {
		sc.Scan()
		h, _ := strconv.Atoi(sc.Text())
		if h >= k {
			ans++
		}
	}
	fmt.Println(ans)
}
