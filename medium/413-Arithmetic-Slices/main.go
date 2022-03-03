package main

func numberOfArithmeticSlices(nums []int) int {
	n, ans, leading := len(nums), 0, 0
	for i := 2; i < n; i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			leading++
			ans += leading
		} else {
			leading = 0
		}
	}
	return ans
}
