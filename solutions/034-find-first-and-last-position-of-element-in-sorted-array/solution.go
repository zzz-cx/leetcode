package main

import "fmt"

// 思路：两次二分——先找第一个等于 target 的下标，再找最后一个
func searchRange(nums []int, target int) []int {
	first := findBound(nums, target, true)
	if first == -1 {
		return []int{-1, -1}
	}
	last := findBound(nums, target, false)
	return []int{first, last}
}

func findBound(nums []int, target int, findFirst bool) int {
	left, right := 0, len(nums)-1
	ans := -1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			ans = mid
			if findFirst {
				right = mid - 1 // 继续向左找更靠前的
			} else {
				left = mid + 1 // 继续向右找更靠后的
			}
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

func main() {
	got := searchRange([]int{5, 7, 7, 8, 8, 10}, 8)
	fmt.Printf("PASS | range=%v\n", got)
}
