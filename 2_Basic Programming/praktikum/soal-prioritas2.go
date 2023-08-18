package main

import "fmt"

func main() {
	segitigaAsterik()
	getFaktorBilangan(26)
}

// SOAL 1
// Buatlah sebuah program untuk mencetak segitiga asterik
func segitigaAsterik() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			if j <= 10/2+i && j >= 10/2-i {
				fmt.Print("*")

			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// SOAL 2
// Buatlah sebuah program untuk mencetak faktor bilangan dari sebuah angka,
func getFaktorBilangan(input int) {
	fmt.Println("Angka", input)
	for i := 1; i <= input; i++ {
		if input%i == 0 {

			fmt.Println(i)
		}
	}
}
