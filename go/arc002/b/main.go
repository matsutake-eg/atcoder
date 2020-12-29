package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	var s string
	fmt.Scan(&s)

	x := strings.Split(s, "/")
	y, _ := strconv.Atoi(x[0])
	m, _ := strconv.Atoi(x[1])
	d, _ := strconv.Atoi(x[2])
	t := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
	for {
		y, m, d := t.Year(), int(t.Month()), t.Day()
		if y%(m*d) == 0 {
			fmt.Printf("%02v/%02v/%02v\n", y, m, d)
			return
		}
		t = t.AddDate(0, 0, 1)
	}
}
