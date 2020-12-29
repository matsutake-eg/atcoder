package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		a int
		b string
	)
	fmt.Scan(&a, &b)
	b = strings.Replace(b, ".", "", -1)
	ib, _ := strconv.Atoi(b)
	fmt.Println(a * ib / 100)
}
