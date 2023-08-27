package main

import "fmt"

type Car struct {
	CarType string
	Fuelin  float64
}

func (c Car) MenghitungPerkiraanBahanBakar() {
	result := float64(c.Fuelin) / 1.5
	fmt.Println("CarType :", c.CarType)
	fmt.Println("est Distance", result)

}

func main() {
	var Car = Car{
		CarType: "SUV",
		Fuelin:  10.5,
	}

	Car.MenghitungPerkiraanBahanBakar()

}
