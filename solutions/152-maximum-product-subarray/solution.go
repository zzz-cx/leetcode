package main

import "fmt"

func maxProduct(nums []int) int {
	maxEnding, minEnding := nums[0], nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxEnding, minEnding = minEnding, maxEnding
		}
		maxEnding = max(nums[i], maxEnding*nums[i])
		minEnding = min(nums[i], minEnding*nums[i])
		ans = max(ans, maxEnding)
	}
	return ans
}

func main() {
	got := maxProduct([]int{2, 3, -2, 4})
	fmt.Printf("PASS | %d\n", got)
}
