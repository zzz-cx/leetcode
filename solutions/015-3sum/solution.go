package main

import (
	"fmt"
	"sort"
)

// threeSum 三数之和（LeetCode 15）
// 思路：排序 + 固定 i + 双指针找两数之和为 -nums[i]
// 易错：找到一组后必须 left++、right--，否则会无限 append 导致 OOM
func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return nil
	}
	sort.Ints(nums)
	var ans [][]int

	for i := 0; i < n-2; i++ { // 至少留两个给 left/right
		if i > 0 && nums[i] == nums[i-1] {
			continue // 跳过重复的 i
		}
		if nums[i] > 0 {
			break // 已排序，后面都 > 0，不可能和为 0
		}
		target := -nums[i]
		left, right := i+1, n-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				// 去重后收缩指针（顺序不能省，最后必须 left++ right--）
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return ans
}

// ThreeSum 兼容旧调用名
func ThreeSum(nums []int) [][]int {
	return threeSum(nums)
}

func main() {
	got := threeSum([]int{-1, 0, 1, 2, -1, -4})
	status := "PASS"
	if len(got) != 2 {
		status = "FAIL"
	}
	fmt.Printf("%s | %d triplets\n", status, len(got))
}
