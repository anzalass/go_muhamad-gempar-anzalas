package controller

func KalkulatorPlus(number ...float64) float64 {
	var total float64
	for _, num := range number {
		total = total + num
	}

	return total
}
func KalkulatorMinus(number ...int) int {
	var total int = number[0]
	for _, num := range number {
		total = total - num
	}

	total = total + number[0]

	return total
}
func KalkulatorKali(x ...int) int {
	var total int = 1
	for _, num := range x {
		total *= num
	}

	return total
}

func KalkulatorBagi(x, y float64) float64 {
	return x / y
}
