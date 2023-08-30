package soaleksplorasi

import (
	"fmt"
	"sort"
)

// SOAL EKSPLORASI NO.1
func MaxSequence(nums []int) int {

	maxSum := nums[0]
	fmt.Println(maxSum, "maxsum")
	currentSum := nums[0]
	fmt.Println(currentSum, "currentSum")

	for i := 1; i < len(nums); i++ {
		fmt.Println(nums[i], "num i")
		if currentSum < 0 {
			currentSum = nums[i]
			// fmt.Println(currentSum, "min")
		} else {
			currentSum = currentSum + nums[i]
			// fmt.Println(currentSum, "+")

		}
		fmt.Println(maxSum, "max")
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	fmt.Println(currentSum, "min")

	return maxSum
}

// SOAL EKSPLORASI NO.2
func MaximumBuyProduct(money int, productPrice []int) int {

	sort.Ints(productPrice)
	fmt.Println(productPrice)
	belanjaan := 0
	totalHargaBelanjaan := 0

	for _, harga := range productPrice {
		if totalHargaBelanjaan+harga <= money {
			totalHargaBelanjaan = totalHargaBelanjaan + harga
			belanjaan++
			fmt.Println("harga  :", harga)
			fmt.Println("belanjaan  :", belanjaan)
			fmt.Println("	total harga belanjaan :", totalHargaBelanjaan)
		}
	}
	fmt.Println(belanjaan)
	return belanjaan
}
