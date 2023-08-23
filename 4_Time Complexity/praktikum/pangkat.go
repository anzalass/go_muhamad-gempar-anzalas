package main

import "fmt"

func main() {
	fmt.Println(hitungPangkat(2, 3)) // 8

	fmt.Println(hitungPangkat(5, 3)) // 125

	fmt.Println(hitungPangkat(10, 2)) // 100

	fmt.Println(hitungPangkat(2, 5)) // 32

	fmt.Println(hitungPangkat(7, 3)) // 343

}

func hitungPangkat(input1 int, input2 int) int {
	pangkat := 1

	for i := 0; i < input2; i++ {
		pangkat = pangkat * input1
	}

	return pangkat

}
