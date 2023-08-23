package main

import "fmt"

func PairSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for _, num := range nums {
		complement := target - num
		if index, found := numMap[num]; found {
			return []int{nums[index], num}
		}
		numMap[complement] = len(numMap)
	}

	return nil
}

func main() {

	// Test cases

	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]

	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11)) // [0, 2]

	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12)) // [2, 3]

	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10)) // [1, 2]

	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6)) // [0, 1]

}
