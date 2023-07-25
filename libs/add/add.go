package add

func Add(args ...int) int {
	sum := 0
	for _, v := range args {
		sum = sum + v
	}
	return sum
}
