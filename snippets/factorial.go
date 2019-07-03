package algo

func fac(x int) int {
	ans := 1
	for i := x; i > 0; i-- {
		ans *= i
	}
	return ans
}
