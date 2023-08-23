package main

import (
	"fmt"
)

func munculSekali(input string) []int {
	charCount := make(map[rune]int)
	var result []int

	for _, char := range input {
		charCount[char]++
	}

	for _, char := range input {
		if charCount[char] == 1 {
			result = append(result, int(char-'0')) // Mengonversi rune ke integer
		}
	}

	return result
}

func main() {

	output := munculSekali("1234123")
	fmt.Println(output) // Output: [4]
}
