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
	for i := 0; i < n; i++ {
		a := scanInt()
		if a%2 == 1 {
			fmt.Println("first")
			return
		}
	}
	fmt.Println("second")
}
