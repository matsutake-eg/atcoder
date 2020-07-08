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
	rc, bc := 0, 0
	for i := 0; i < n; i++ {
		s := scanString()
		for _, r := range s {
			switch r {
			case 'R':
				rc++
			case 'B':
				bc++
			}
		}
	}
	if rc == bc {
		fmt.Println("DRAW")
	} else if rc > bc {
		fmt.Println("TAKAHASHI")
	} else {
		fmt.Println("AOKI")
	}
}
