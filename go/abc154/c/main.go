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
	n := readInt()
	m := make(map[int]bool)
	for i := 0; i < n; i++ {
		a := readInt()
		m[a] = true
	}

	if n == len(m) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
