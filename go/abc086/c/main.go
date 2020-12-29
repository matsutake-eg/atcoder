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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	n := readInt()
	bt, bx, by := 0, 0, 0
	for i := 0; i < n; i++ {
		t, x, y := readInt(), readInt(), readInt()
		dt, dx, dy := t-bt, abs(x-bx), abs(y-by)
		if mv := dx + dy; mv > dt || dt%2 != mv%2 {
			fmt.Println("No")
			return
		}
		bt, bx, by = t, x, y
	}
	fmt.Println("Yes")
}
