package main

import (
	"fmt"
	"strings"
)

func checkPalindrom(input string) bool {
	input = strings.ToLower(strings.ReplaceAll(input, " ", ""))

	for i := 0; i < len(input)/2; i++ {
		if input[i] == input[len(input)-1-i] {
			return true
		}
	}
	return false

}

func main() {
	var input string

	fmt.Print("Masukkan kata atau frasa: ")
	fmt.Scan(&input)
	if checkPalindrom(input) {
		fmt.Println(input, "Adalah Palindrom")
	} else {
		fmt.Println(input, "Bukan Palindrom")
	}
}
