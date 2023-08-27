package main

import (
	"fmt"
	"math"
)

type Student struct {
	Name  string
	Score float64
}

type StudentList []Student

func (stdlist StudentList) NilaiMaksimal() Student {
	nilaiMax := -math.MaxFloat64
	var nilaiMaxStd Student

	for _, std := range stdlist {
		if std.Score > nilaiMax {
			nilaiMax = std.Score
			nilaiMaxStd = std

		}

	}

	return nilaiMaxStd

}
func (stdlist StudentList) NilaiMinimal() Student {
	nilaiMin := math.MaxFloat64
	var nilaiMinStd Student

	for _, std := range stdlist {
		if std.Score < nilaiMin {
			nilaiMin = std.Score
			nilaiMinStd = std

		}

	}

	return nilaiMinStd

}

func (stdlist StudentList) RataRata() float64 {
	nilai := 0.0

	for _, std := range stdlist {
		nilai = nilai + std.Score
	}

	avg := nilai / float64(len(stdlist))
	return avg

}

func main() {
	daftarSiswa := make(StudentList, 0)

	for i := 1; i < 6; i++ {
		var name string
		var score float64
		fmt.Printf("Masukan nama siswa ke-%d", i)
		fmt.Scanln(&name)
		fmt.Printf("Masukan ni;ai siswa ke-%d", i)
		fmt.Scanln(&score)

		daftarSiswa = append(daftarSiswa, Student{Name: name, Score: score})
	}

	avg := daftarSiswa.RataRata()
	max := daftarSiswa.NilaiMaksimal()
	min := daftarSiswa.NilaiMinimal()

	fmt.Printf("Skor rata-rata: %.2f\n", avg)
	fmt.Printf("Skor Tertinggi adalah : %s, dengan nilai %.2f\n", max.Name, max.Score)
	fmt.Printf("Skor Terendah adalah : %s, dengan nilai %.2f\n", min.Name, min.Score)
}
