package main

import (
	"fmt"
	"math"
)

func getMinMax(numbers ...*int) (min, max int) {

	// your code here
	min = math.MaxInt64
	max = -math.MaxInt64

	for _, v := range numbers {
		num := *v
		if num < min {
			min = num
		}

	}

	for _, v := range numbers {
		numM := *v
		if numM > max {
			max = numM
		}
	}

	return min, max

}

func main() {

	var a1, a2, a3, a4, a5, a6, min, max int

	fmt.Scan(&a1)

	fmt.Scan(&a2)

	fmt.Scan(&a3)

	fmt.Scan(&a4)

	fmt.Scan(&a5)

	fmt.Scan(&a6)

	fmt.Println()

	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)

	fmt.Println("Nilai Min : ", min)

	fmt.Println("Nilai Max : ", max)

}
