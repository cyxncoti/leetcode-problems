package main

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if nums[1] < nums[0] {
		nums[1] = nums[0]
	}
	for i := 2; i < n; i++ {
		if rob := nums[i] + nums[i-2]; rob > nums[i-1] {
			nums[i] = rob
		} else {
			nums[i] = nums[i-1]
		}
	}
	return nums[n-1]
}
