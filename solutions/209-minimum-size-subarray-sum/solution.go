package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	left := 0
	total := 0
	ans := len(nums) + 1

	for right := 0; right < len(nums); right++ {
		total += nums[right]
		for total >= target {
			if right-left+1 < ans {
				ans = right - left + 1
			}
			total -= nums[left]
			left++
		}
	}

	if ans == len(nums)+1 {
		return 0
	}
	return ans
}

func main() {
	tests := []struct {
		target   int
		nums       []int
		expected int
	}{
		{7, []int{2, 3, 1, 2, 4, 3}, 2},
		{4, []int{1, 4, 4}, 1},
		{11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},
	}

	for _, tc := range tests {
		result := minSubArrayLen(tc.target, tc.nums)
		status := "PASS"
		if result != tc.expected {
			status = "FAIL"
		}
		fmt.Printf("%s | target=%d, nums=%v => %d (expected %d)\n", status, tc.target, tc.nums, result, tc.expected)
	}
}
