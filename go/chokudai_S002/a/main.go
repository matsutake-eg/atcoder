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
	var a, b int
	for i := 0; i < n; i++ {
		sc.Scan()
		a, _ = strconv.Atoi(sc.Text())
		sc.Scan()
		b, _ = strconv.Atoi(sc.Text())
		fmt.Println(a * b)
	}
}
