package main

import (
	"bufio"
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

func isPrime(x int) bool {
	if x < 2 {
		return false
	}

	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	isp := make([]bool, 100001)
	for i := 1; i < len(isp); i++ {
		isp[i] = isPrime(i)
	}
	s := make([]int, 100001)
	for i := 1; i < len(s); i++ {
		lk := 0
		if isp[i] && isp[(i+1)/2] {
			lk = 1
		}
		s[i] = lk + s[i-1]
	}

	var wr = bufio.NewWriter(os.Stdout)
	q := readInt()
	for i := 0; i < q; i++ {
		l, r := readInt(), readInt()
		wr.WriteString(strconv.Itoa(s[r]-s[l-1]) + "\n")
	}
	wr.Flush()
}
