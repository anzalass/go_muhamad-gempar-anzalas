package main

import (
	"fmt"
	"math"
)

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}

	diagonal1 := 0
	diagonal2 := 0
	size := len(matrix)

	for i := 0; i < size; i++ {
		diagonal1 += matrix[i][i]
	}

	for i := 0; i < size; i++ {
		diagonal2 += matrix[i][size-i-1]
	}

	fmt.Println(size)
	fmt.Println(diagonal1)
	fmt.Println(diagonal2)

	Selisih := int(math.Abs(float64(diagonal1 - diagonal2)))

	fmt.Println("Selisih absolut antara jumlah diagonal:", Selisih)
}
