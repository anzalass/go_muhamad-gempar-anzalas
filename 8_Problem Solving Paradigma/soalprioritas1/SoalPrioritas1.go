package soalprioritas1

import (
	"fmt"
	"math"
)

func CariBiner(n int) []int {
	ans := make([]int, n+1)

	for i := 1; i <= n; i++ {
		ans[i] = ans[i-1] << 1
		ans[i] += 1
	}

	return ans
}

func Pascal(numRows int) [][]int {

	ans := make([][]int, numRows)
	for i := range ans {
		ans[i] = make([]int, i+1)
	}

	ans[0][0] = 1
	for i := 1; i < numRows; i++ {
		ans[i][0] = 1
		ans[i][i] = 1
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
		}
	}

	return ans
}

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

func SimpleEquations(a, b, c float64) {

	if b*b-4*a*c < 0 {
		fmt.Println("No solution")
		return
	}

	x1 := (-b + math.Sqrt(b*b-4*a*c)) / (2 * a)
	x2 := (-b - math.Sqrt(b*b-4*a*c)) / (2 * a)

	fmt.Println(x1, x2)
}
