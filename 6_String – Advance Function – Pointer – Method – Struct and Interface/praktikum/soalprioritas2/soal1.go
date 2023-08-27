package main

import "fmt"

// func caesar(offset int, input string) string {

// 	arr := make([]byte, 0)

// 	// your code here
// 	for i := 0; i < len(input); i++ {
// 		tes1 := (input[i] + byte(offset)) % 26
// 		tes1 = tes1 + 'a'
// 		arr = append(arr, tes1)
// 	}
// 	return string(arr)

// }

func caesar(offset int, input string) string {
	arr := make([]byte, 0)

	for i := 0; i < len(input); i++ {
		char := input[i]
		if char >= 'a' && char <= 'z' {
			tes1 := ((char - 'a') + byte(offset)) % 26
			tes1 += 'a'
			arr = append(arr, tes1)
		}
	}
	return string(arr)
}

func main() {

	fmt.Println(caesar(3, "abc")) // def

	fmt.Println(caesar(2, "alta")) // def

	fmt.Println(caesar(10, "alterraacademy")) // kvdobbkkmknowi

	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) // bcdefghijklmnopqrstuvwxyza

	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl

}
