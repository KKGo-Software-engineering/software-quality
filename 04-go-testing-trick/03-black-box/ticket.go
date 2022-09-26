package ticket

func Price(age int) float64 {
	if age <= 3 {
		return 0
	}
	if age <= 15 {
		return 15
	}
	if age <= 50 {
		return 30
	}
	return 5
}
