package main

import (
	"bufio"
	"fmt"
	"math"
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
	txa, tya := scanInt(), scanInt()
	txb, tyb := scanInt(), scanInt()
	t, v := scanInt(), scanInt()
	n := scanInt()
	for i := 0; i < n; i++ {
		x, y := scanInt(), scanInt()
		dxa, dya := x-txa, y-tya
		dxb, dyb := x-txb, y-tyb
		if math.Sqrt(float64(dxa*dxa+dya*dya))+math.Sqrt(float64(dxb*dxb+dyb*dyb)) <= float64(t*v) {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
}
