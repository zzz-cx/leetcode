package main

import "fmt"

// 思路：每次 mid 必有一侧仍是有序段，先判断 target 是否落在该有序段内，再缩区间
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] { // 左半段有序
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // 右半段有序
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func main() {
	got := search([]int{4, 5, 6, 7, 0, 1, 2}, 0)
	status := "PASS"
	if got != 4 {
		status = "FAIL"
	}
	fmt.Printf("%s | index=%d\n", status, got)
}
