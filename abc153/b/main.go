package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	iv, _ := strconv.Atoi(sc.Text())
	return iv
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	h, n := readInt(), readInt()
	sum := 0
	for i := 0; i < n; i++ {
		a := readInt()
		sum += a
	}
	if sum >= h {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
