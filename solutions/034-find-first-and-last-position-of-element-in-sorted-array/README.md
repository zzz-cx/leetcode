# 在排序数组中查找元素的第一个和最后一个位置

> LeetCode 34 · [find-first-and-last-position-of-element-in-sorted-array](https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/)

## 题目

给定升序数组 `nums` 和目标值 `target`，找出目标值在数组中的开始位置和结束位置。

## 题解思路与解析

- 思路：两次二分——先找第一个等于 target 的下标，再找最后一个

## 解答

### Golang

```go
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
```

### Python

```python
# 思路：两次二分——先找第一个等于 target 的下标，再找最后一个
from typing import List


class Solution:
    def search_range(self, nums: List[int], target: int) -> List[int]:
        first = self._find_bound(nums, target, True)
        if first == -1:
            return [-1, -1]
        last = self._find_bound(nums, target, False)
        return [first, last]

    def _find_bound(self, nums: List[int], target: int, find_first: bool) -> int:
        left, right = 0, len(nums) - 1
        ans = -1
        while left <= right:
            mid = (left + right) // 2
            if nums[mid] == target:
                ans = mid
                if find_first:
                    right = mid - 1
                else:
                    left = mid + 1
            elif nums[mid] < target:
                left = mid + 1
            else:
                right = mid - 1
        return ans
```
