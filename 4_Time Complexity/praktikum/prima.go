package main

import "fmt"

func main() {
	fmt.Println(checkPrima(1000000007)) // true

	fmt.Println(checkPrima(13)) // true

	fmt.Println(checkPrima(17)) // true

	fmt.Println(checkPrima(20)) // false

	fmt.Println(checkPrima(35)) // false
}

func checkPrima(input int) bool {
	ctr := 0
	for i := 1; i <= input; i++ {
		if input%i == 0 {
			ctr++
		}
	}

	return ctr == 2

}
