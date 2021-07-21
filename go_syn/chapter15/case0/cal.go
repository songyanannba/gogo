package cal

func AddUpper(n int) int {
	res := 0
	for i := 1; i <= n-1; i++ {
		res += i
	}
	return res
}

func GetSub(n1, n2 int) int {
	return n1 - n2
}
