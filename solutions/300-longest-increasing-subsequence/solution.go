package main

import "fmt"

func lengthOfLIS(nums []int) int { //贪心算法+二分查找解法
	if len(nums) == 0 {
		return 0
	}
	d := make([]int, len(nums)+1)
	lens := 1
	d[lens] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > d[lens] { //如果当前元素大于d[len]，则将当前元素加入d中
			d[lens+1] = nums[i]
			lens++
		} else { //如果当前元素小于等于d[len]，则二分查找d中第一个大于等于当前元素的位置
			pos := binarySearch(d, 1, lens, nums[i])
			d[pos] = nums[i]
		}
	}
	return lens
}
func binarySearch(d []int, left, right, target int) int {
	for left < right {
		mid := (left + right) / 2
		if d[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
func lengthOfLIS2(nums []int) int { //动态规划解法
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	ans := 0
	for _, v := range dp {
		ans = max(ans, v)
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	got := lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})
	fmt.Printf("PASS | %d\n", got)
}
