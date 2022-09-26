package sum

func sum(xs ...int) int {
	var total int
	for _, num := range xs {
		total += num
	}
	return total
}
