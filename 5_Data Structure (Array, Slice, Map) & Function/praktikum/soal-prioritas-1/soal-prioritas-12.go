package main

import "fmt"

func Mapping(input []string) map[string]int {
	res := make(map[string]int)

	for _, item := range input {
		res[item]++
	}

	return res
}

func main() {
	res := Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"})
	fmt.Println(res)
}
