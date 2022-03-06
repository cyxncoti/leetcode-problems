package main

import "fmt"

func main() {
	test(1, 1)
	test(2, 6)
	test(3, 90)
	test(7, 681080400)
	test(8, 729647433)
	test(9, 636056472)
}

func test(input int, answer int) {
	output := countOrders(input)
	if output == answer {
		return
	}
	fmt.Printf("[Fail] Input = %d, output = %d, answer = %d\n", input, output, answer)
}

func countOrders(n int) int {
	ans, mod := 1, 1000000007
	for i := 2; i <= 2*n; i++ {
		ans *= i
		if i%2 == 0 {
			ans /= 2
		}
		ans %= mod
	}
	return ans
}
