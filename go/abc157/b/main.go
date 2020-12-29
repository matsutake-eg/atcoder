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
	a := make([][]int, 3)
	for i := range a {
		a[i] = make([]int, 3)
		for j := range a[i] {
			a[i][j] = scanInt()
		}
	}

	n := scanInt()
	for i := 0; i < n; i++ {
		b := scanInt()
		for i := range a {
			for j := range a[i] {
				if a[i][j] == b {
					a[i][j] = 0
				}
			}
		}
	}

	if a[0][0] == 0 && a[0][1] == 0 && a[0][2] == 0 {
		fmt.Println("Yes")
	} else if a[1][0] == 0 && a[1][1] == 0 && a[1][2] == 0 {
		fmt.Println("Yes")
	} else if a[2][0] == 0 && a[2][1] == 0 && a[2][2] == 0 {
		fmt.Println("Yes")
	} else if a[0][0] == 0 && a[1][0] == 0 && a[2][0] == 0 {
		fmt.Println("Yes")
	} else if a[0][1] == 0 && a[1][1] == 0 && a[2][1] == 0 {
		fmt.Println("Yes")
	} else if a[0][2] == 0 && a[1][2] == 0 && a[2][2] == 0 {
		fmt.Println("Yes")
	} else if a[0][0] == 0 && a[1][1] == 0 && a[2][2] == 0 {
		fmt.Println("Yes")
	} else if a[0][2] == 0 && a[1][1] == 0 && a[2][0] == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
