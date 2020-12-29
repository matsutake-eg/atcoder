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

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	d := make([]int, 7)
	for i := range d {
		d[i] = scanInt()
	}
	j := make([]int, 7)
	for i := range j {
		j[i] = scanInt()
	}

	ans := 0
	for i := 0; i < 7; i++ {
		ans += max(d[i], j[i])
	}
	fmt.Println(ans)
}
