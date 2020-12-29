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
	for i := 0; i < n; i++ {
		a := readInt()
		if a%2 == 0 {
			if !(a%3 == 0 || a%5 == 0) {
				fmt.Println("DENIED")
				return
			}
		}
	}
	fmt.Println("APPROVED")
}
