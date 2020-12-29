package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	iv, _ := strconv.Atoi(scanString())
	return iv
}

func scanString() string {
	sc.Scan()
	return sc.Text()
}

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 100000), 100000000)
}

func main() {
	n := scanInt()
	a, b, c, d, e, f := 0, 0, 0, 0, 0, 0
	for i := 0; i < n; i++ {
		M, _ := strconv.ParseFloat(scanString(), 64)
		m, _ := strconv.ParseFloat(scanString(), 64)
		if M >= 35.0 {
			a++
		} else if M >= 30.0 {
			b++
		} else if M >= 25.0 {
			c++
		}
		if m >= 25.0 {
			d++
		} else if m < 0.0 && M >= 0.0 {
			e++
		} else if M < 0.0 {
			f++
		}
	}
	fmt.Println(a, b, c, d, e, f)
}
