package main

import "fmt"

func main() {

	// masukan alas_bawah di parameter pertama
	// masukan alas_atas di parameter kedua
	// masukan tinggi di parameter ketiga
	fmt.Println(luasTrapesium(16, 10, 8))

	// masukan angka yang ingin di check pada parameter
	checkGanjilGenap(98)

	// masukan nilai yang ingin di check pada parameter
	fmt.Println(checkNilaidanGrade(888))

	loopFizzBuzz()
}

// SOAL NO 1
// Buatlah sebuah program untuk menghitung luas trapesium.
func luasTrapesium(alas_bawah int, alas_atas int, tinggi int) int {

	hasilHitung := (alas_bawah + alas_atas) / 2 * tinggi
	return hasilHitung

}

// SOAL NO 2
// Buatlah sebuah program untuk menentukan apakah sebuah bilang adalah bilang ganjil atau genap.
func checkGanjilGenap(angka int) {
	if angka%2 == 0 {
		fmt.Println(angka, "Adalah Angka Genap")
	} else {
		fmt.Println(angka, "Adalah Angka Ganjil")
	}
}

// SOAL NO 3
// Buatlah sebuah program untuk menentukan grade dari sebuah nilai, dengan ketentuan sebagai berikut:
func checkNilaidanGrade(nilai int) string {

	if nilai > 80 && nilai < 100 {
		return "A"
	} else if nilai > 65 && nilai < 79 {
		return "B"
	} else if nilai > 50 && nilai < 64 {
		return "C"
	} else if nilai > 35 && nilai < 49 {
		return "D"
	} else if nilai > 0 && nilai < 34 {
		return "E"
	} else {
		return "Nilai Invalid"
	}

}

// SOAL NO 4
// Buatlah sebuah program yang mencetak angka dari 1 sampai 100 dan untuk kelipatan '3' cetak "Fizz" sebagai ganti angka, dan untuk kelipatan '5' cetak "Buzz”. Sebagai contoh; 1 2 fizz 4 buzz fizz 7 8 fizz buzz …….
func loopFizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
