package main

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	var w int
	fmt.Scan(&w)

	s := "DiscoPresentsDiscoveryChannelProgrammingContest2016"
	for i := 0; i < len(s); i += w {
		fmt.Println(s[i:min(len(s), i+w)])
	}
}
