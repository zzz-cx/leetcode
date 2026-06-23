package main

import "fmt"

func maxSubArray(nums []int) int { //思路，只要前面的数字和大于零，就把他加上，如果小于0，则当前就是最大的，重新计算

	n := len(nums)
	maxSum := nums[0]
	for i := 1; i < n; i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if nums[i] > maxSum {
			maxSum = nums[i]
		}
	}
	return maxSum
}

func main() {
	got := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Printf("PASS | sum=%d\n", got)
}
