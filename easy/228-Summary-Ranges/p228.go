package main

import "strconv"

func summaryRanges(nums []int) []string {
	output := []string{}
	for i, n := 0, len(nums); i < n; i++ {
		start := nums[i]
		for ; i+1 < n && nums[i+1] == nums[i]+1; i++ {
		}
		if nums[i] == start {
			output = append(output, strconv.Itoa(start))
		} else {
			output = append(output, strconv.Itoa(start)+"->"+strconv.Itoa(nums[i]))
		}
	}
	return output
}
