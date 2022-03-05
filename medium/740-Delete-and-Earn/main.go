package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	input := []int{3, 4, 2}
	fmt.Printf("Input = %v, output = %d, answer = 6\n", input, deleteAndEarn(input))
	input = []int{2, 2, 3, 3, 3, 4}
	fmt.Printf("Input = %v, output = %d, answer = 9\n", input, deleteAndEarn(input))
	input = []int{1}
	fmt.Printf("Input = %v, output = %d, answer = 1\n", input, deleteAndEarn(input))
}

func deleteAndEarn(nums []int) int {
	points, maxNumber := map[int]int{}, 0
	for _, v := range nums {
		points[v] += v
		if maxNumber < v {
			maxNumber = v
		}
	}
	n, twoBack, oneBack := len(points), 0, 0
	if float64(maxNumber) < float64(n)+float64(n)*math.Log2(float64(n)) {
		oneBack = points[1]
		for i := 2; i <= maxNumber; i++ {
			take := points[i] + twoBack
			twoBack = oneBack
			if take > oneBack {
				oneBack = take
			}
		}
	} else {
		values := make([]int, n)
		i := 0
		for v := range points {
			values[i] = v
			i++
		}
		sort.Ints(values)
		oneBack = points[values[0]]
		for i = 1; i < n; i++ {
			if values[i]-values[i-1] > 1 {
				twoBack = oneBack
				oneBack += points[values[i]]
				continue
			}
			take := points[values[i]] + twoBack
			twoBack = oneBack
			if take > oneBack {
				oneBack = take
			}
		}
	}
	return oneBack
}
