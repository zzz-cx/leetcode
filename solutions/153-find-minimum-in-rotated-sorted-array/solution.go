package main

import "fmt"

// 思路：二分缩区间，用 nums[mid] 与 nums[right] 比较判断最小值在左半还是右半
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

func main() {
	got := findMin([]int{3, 4, 5, 1, 2})
	fmt.Printf("PASS | %d\n", got)
}
