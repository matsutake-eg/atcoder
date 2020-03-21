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
	cnt4 := 0
	cnt2 := 0
	for i := 0; i < n; i++ {
		a := scanInt()
		if a%4 == 0 {
			cnt4++
		} else if a%2 == 0 {
			cnt2++
		}
	}

	if n%2 == 1 {
		n--
	}
	n -= cnt4 * 2
	n -= cnt2
	if n <= 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
