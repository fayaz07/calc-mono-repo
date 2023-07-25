package subtract

func Subtract(args ...int) int {
	res := 0
	for _, v := range args {
		res = res - v
	}
	return res
}
