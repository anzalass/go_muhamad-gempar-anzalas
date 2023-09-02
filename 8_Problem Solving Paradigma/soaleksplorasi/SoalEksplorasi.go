package soaleksplorasi

import "fmt"

func Romawi(n int) string {

	simbol := []string{"M", "D", "C", "L", "X", "V", "I"}
	values := []int{1000, 500, 100, 50, 10, 5, 1}

	output := ""

	for i := 0; n > 0; i++ {
		count := n / values[i]
		fmt.Println(n, values[i])
		fmt.Println(count)

		for j := 0; j < count; j++ {
			output = output + simbol[i]
		}

		n -= count * values[i]
		fmt.Println(n)
	}

	return output
}
