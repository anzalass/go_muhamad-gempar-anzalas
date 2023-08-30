package soalprioritas1

import (
	_ "fmt"
	"sort"
)

// SOAL PRIOSITAS 1 NO.1
func Fibonacci(number int) int {
	if number < 1 {
		return 0
	}

	var fib = make([]int, number+1)
	fib[0] = 0
	fib[1] = 1

	for i := 2; i <= number; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib[number]

}

// SOAL PRIORITAS 1 NO.2
type pair struct {
	name  string
	count int
}
type PairSlice []pair

func (p PairSlice) Len() int {
	return len(p)
}
func (p PairSlice) Less(i, j int) bool {
	return p[i].count < p[j].count
}
func (p PairSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func MostAppearItem(items []string) []pair {
	countMap := make(map[string]int)

	for _, item := range items {
		countMap[item]++
	}

	pair2 := make([]pair, 0, len(items))
	for idx, value := range countMap {
		pair2 = append(pair2, pair{idx, value})
	}

	return pair2
}

// SOAL PRIORITAS 1 NO.3
func getPrime(n int) bool {
	var ipt = 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			ipt++
		}
	}
	return ipt == 2
}

func PrimeX(number int) int {

	// your code here
	map1 := make(map[int]int)
	map2 := make(map[int]int)

	for i := 0; i <= 100; i++ {
		map1[i]++
	}

	for i := 1; i <= 1000; i++ {
		if getPrime(i) == true {
			map2[i]++
		}
	}

	arr := make([]int, 0, len(map2))
	for idx, _ := range map2 {
		arr = append(arr, idx)
	}

	sort.Ints(arr)
	return arr[number-1]

}
